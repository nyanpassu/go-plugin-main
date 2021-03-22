.PHONY: build install
.DEFAULT: default

build:
	go build -buildmode=plugin -o systemd-task-linux-amd64.so plugin.go

clean:
	rm systemd-task-linux-amd64.so

install:
	cp systemd-task-linux-amd64.so /var/lib/containerd/plugins/

uninstall:
	rm /var/lib/containerd/plugins/systemd-task-linux-amd64.so
