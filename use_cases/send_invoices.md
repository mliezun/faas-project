## Sending invoices

Generate invoices from md files and send them via email.

Example invoice:

```md
# Invoice

Date: 2023-04-01

Due date: 2023-04-10

Charges for services given during the month of March

| Date       | Description | Qty | Price           | Total |
|------------|-------------|-----|-----------------|-------|
| 2023-03-01 | Web service | 1   | $ 10            | $ 10  |
|            |             |     | **Grand Total** | $ 10  |

### To be paid: $ 10
```

See result [here](/use_cases/invoice.pdf). Obtained using pandoc.



### Backend code

In `main.fs` file define the functions to be executed:

```
fn generate_pdf:
    markdown = "..."                    # Inline string or
    markdown = read_file("template.md") # read from local file

    mardown.render(
        date=date.today()
    )

    p = get_plugin("md_to_pdf", "v1")   # Version is optional, default: latest
    p.generate(markdown)                # Last expression is the result

fn email_to_customer(file):
    customer = "customer@example.com"

    e = get_plugin("email")
    e.send(                             # Send email to customer with pdf attached
        customer,
        "subject",
        "body",
        file,
    )

# This creates a connection between the two functions
send_pdf = generate_pdf -> email_to_customer
```

In `schedule.fs` you can set crons for functions or connections.
Inside this file you have access to anything that was declared in `main.fs`.

```
Every first day of the month: send_pdf  # Schedule using natural language
0 0 1 * *: send_pdf                     # Schedule using cron syntax
```


Then execute the cli tool to upload to server:

```
$ cli-tool sync
```

### Testing

Ideally use the CLI tool for testing.

#### Examples

Failing email

```
$ cli-tool test send_pdf
...
Calling generate_pdf            OK
    See file generated here: https://fs.com/assets/<uuid>
        - Expires after 5min
Calling email_to_customer       FAILED
    Plugin "email" is not enabled
        - Setup email integration first
        - Go to website docs: https://fs.com/docs/email-integration
```

Success email

```
$ cli-tool test send_pdf
...
Calling generate_pdf            OK
    See file generated here: https://fs.com/assets/<uuid>
        - Expires after 5min
Calling email_to_customer       OK
    See email sent here: https://fs.com/emails/<uuid>
        - Expires after 5min
```


## TODO

- Programmatic testing were outputs can be checked by code, or maybe AI?
