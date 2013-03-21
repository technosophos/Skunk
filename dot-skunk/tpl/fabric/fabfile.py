from fabric.api import *

@task
def whoami():
	puts("{{.project}} by {{.author}}.")
