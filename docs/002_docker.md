# Linux Containers

## What are LXC

[namespace](https://man7.org/linux/man-pages/man7/namespaces.7.html) is Linux feature that partition kernel resources such that one set of
processes sees one set of resources. An abstraction that makes it appear to the
processes to the namespace that they have their own isolated instance of the
global resources. Changes to that global resources are visible to the processes
only on that namesapce, but invisible to the other processes.

[cgroups](https://man7.org/linux/man-pages/man7/cgroups.7.html) (control groups) is a Linux kernel feature that limits, accounts for,
and isolates the resource usage(CPU, memory, disk I/O) of a collection of
processes. They organize processes into hierarchy groups whose usage of
resources can limited and monitored.

[unionfs](https://en.wikipedia.org/wiki/UnionFS) is a filesystem service that
implement union mounting (a way of combining multiple direcotries into one that
appears to contain their combined contents). It allows files and dirs to be
transparently overlay. Docker uses this to implement its layer images.

seccomp-bpf (secure computing mode) filter system calls that are available to a process

## Benefits of using containers

Easier to apply Devops pracitce such as:

- Deployment time is significantly faster;<br>
- Faster to build that a VM image;<br>
- Easier and fast to scale;<br>
- Run and test the same iamge across all environements;<br>
- Same deployment processes across environements;<br>
- You need to maintain only the libraries and the runtime;<br>
- Better resource utilization;<br>
- Increase productivity;<br>
- Encourages service oriented architecture;<br>

## Specs

dockerfiles

image

container

container networking drivers:

- bridge (default);<br>
- Host - remove network isolation between host and container;<br>
- overlay - connects multiple container deamons together;<br>
- macvlan - assign mac address to the container so the container to appear as physical device on the network;<br>
- none;<br>

## Best Practices

- use multistage builds
- build within `builder` run in `final`
- use `user` in final image
- expose port within the dockerimage itself (it's nice practice)
- optional usage of CMD and ENTRYPOINT:

``` dockerfile
CMD["./cmd_to_run", "arg1"] # default command to run the container
ENTRYPOINT [""] # optional; if passed arguments it will take precedence
```

## Additional Resources

https://dockerbook.com

https://docs.docker.com
