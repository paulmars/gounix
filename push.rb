require 'pusher'

Pusher.url = "https://bb9289e6521feadf6999:1e97d762ff900c953e06@api.pusherapp.com/apps/130573"

Pusher['test_channel'].trigger('my_event', {
  message: 'hello world'
})

