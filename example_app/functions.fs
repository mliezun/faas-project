// App deployed to: https://abc123.faast.com
// Queue exposed https://queue.abc123.faast.com/external-trigger?Signature=...

function verify_signature (sig) {
    return true
}

event "http" "/webhook" {
    let data = request.body.json()
    let sample_ids = data.sample_ids
    if verify_signature(request.header["Gencove-Signature"]) {
        queue.process_samples.publish(sample_ids)
        request.write(200, "OK")
    } else {
        request.write(400, "Bad request")
    }
}

event "http" "/samples?project_id=<project_id>" {
    db.get("samples", {"project_id": project_id})
}

event "queue" "process_samples" {
    for sample_id in sample_ids {
        let result = requests.get("https://api.gencove.com/api/v2/sample-detail/${sample_id}")
        db.store("samples", sample_id, result, overwrite=true)
    }
}
