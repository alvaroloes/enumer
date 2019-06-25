workflow "Test on Push" {
  on = "push"
  resolves = ["test"]
}

# Deploy, and write deployment to file
action "test" {
  runs = "make"
  args = "test"
}
