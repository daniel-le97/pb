/// <reference path="../pb_data/types.d.ts" />

routerAdd("POST", "/webhooks/:id", (c) => {
    let name = c.pathParam("id")
   let eventPayload =  $apis.requestInfo(c)
   let headers = eventPayload.headers
    console.log("event payload", eventPayload.data.repository.html_url)
    console.log("headers", headers["x-github-event"])
   
    const githubEvent = new DynamicModel({
        event: '',
        payload: {
            repository: {
                html_url: ''
            }
        },
        pusher:{
            id: 0
        },
        sender:{
            login: ''
        }
    })
    c.bind(githubEvent)
    console.log("github event", githubEvent);

    return c.json(200, { "message": "Hello " + name })
})

// onAfterBootstrap((e) => {
//     const cmd = $os.cmd('bun', 'smee.ts')
// cmd.start()
// // execute the command and return its standard output as string

// })

onModelAfterUpdate((e) => {
    console.log("user updated...", e.model.get("email"))
}, "users")