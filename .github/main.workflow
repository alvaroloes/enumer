workflow "Test on Push" {
  on = "push"
  resolves = ["test"]
}

action "test" {
  uses = "cedrickring/golang-action@1.3.0"
  runs = "make"
  args = "test"
}
