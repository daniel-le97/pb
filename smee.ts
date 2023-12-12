import SmeeClient from 'smee-client'

const smee = new SmeeClient({
  source: 'https://smee.io/2CheYVHetZe4ROm',
  target: 'http://localhost:8090/webhooks/1',
  logger: console
})

const events = smee.start()

// Stop forwarding events
// events.close()