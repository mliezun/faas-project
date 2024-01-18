# Dreaming about the coding of tomorrow
There will always be servers.

Why programming feels so painful?
- Handle dependencies
- Vulnerability scans
- Integration of multiple systems and languages
- Code is hard to explain and understand and logic is split in various places
- Is difficult to design performant code
- Manage permissions (which service has access to what)
- Manage infrastructure: how many servers, in which network, create/destroy resources
- Almost everybody is writting code to do very similar things
- We use programming languages that are wasteful in energy and resources
- Now the framework has async support and we need to change our codebase to support it
- New major version of the programming language (or framework) is out with lots of breaking changes (Python 2 -> 3 transition)
- [The elements of a good information system](https://github.com/fpereiro/elements)


It would be so good to have a consistent way of doing this. That's clear and unified.

How would that look like?

Let's see some examples:

### Copy from S3 to GCP Object Storage
This is a very painful task. I start to imagine all that has to be done and Im crying already.

What needs to be done to acomplish this is:
1. Have AWS and GCP accounts setup
2. Identify the locations where files should be copied to/from
3. Create an AWS Role and GCP service account, where the first one has a permissions to read from a location and the second to write to the destination.
4. Choose a way to implement this: AWS Lambda, GCP Lambda, Microservice running in Kubernetes, Task in a Job Queue, ...
5. Then choose a language to implement this in: Python, Node, Go, Rust, PHP, Java, Scala, Kotlin, Dart, Haskell, Clojure, ...
6. So many options!!
7. Choose libraries to connect to AWS and GCP: infinite posibilities
8. Make sure the credentials to make the copy are available and can be accessed by the program that needs to implement this feature
9. Write the code. No problem here. This is the good part :-)
10. Setup CI/CD, with vulnerability scanning and keep dependencies up-to-date or the deploy will fail
11. Actually use this feature and make sure everything is working as intended

So much to do. In an ideal world, we would only need to do 1, 2, 3 (sad but unavoidable for security), 9 and 11. Eliminating most of the pain from this.



### Licensing

https://www.elastic.co/licensing/elastic-license
