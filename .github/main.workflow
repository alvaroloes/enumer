workflow "Test on Push" {
  on = "push"
  resolves = ["test"]
}

action "test" {
  uses = "actions/action-builder/shell@master"
  runs = "make"
  args = "test"
}
