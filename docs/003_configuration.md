# Configuration Management

**Configuration management** is the systematic management of software,
hardware, and system configurations to ensure consistency and reliability. Why: <br>
- Consistency: ensures all systems are in a known and a consistent state
- Traceability: track changes and identifies their impact
- Efficiency: streamlines processes and reduces downtime
- Complience: Ensure adherence to industry regulation and standarts
- Reproducibility: Enables the recreation of specific configuration

**Idempotence** is a fundamental concept in computer science that means
performing an operation multiple times has the same effect as performing it
once. It gives: <br>
- Predicatbility
- Repeatability
- Error Handling
- Efficiency

**Best practices**: <br>
- Documentation: Thorougly document configurations and changes
- Change Control: Establish and follow a structured change control process
- Automation: Automate repetitive tasks for efficiency
- Modularity: Organize configurations into modular components
- Regular auditing: Periodically audit configurations for accuracy
- Training: Train and educate staff on best practices

**Tools**: <br>
- Ansible - agentless; using ssh to connect and run playbooks; write playbooks using plain YAML;
- CHEF - Uses pure Ruby; Requires agents installed on the machines (acquired by Progress)
- Puppet - Requires an agent installed on the machines; includes its own
  declaratice to describe system configuration or Ruby;

**Why Ansible**: <br>
- Simplicity: yaml
- Agentless
- Idempotence
- Scalability: manage thousands of systems in parallel
- Community and Ecosystem: Large and active user community

**Primitives**:<br>
- Task: single unit of work (install a pkg, copy file, etc) within a play
- Playbook: a YAML file that defines a set of plays, which are series of tasks
  executed on a host. Playbooks serves as top-level structure for Ansible
  automation and allow you to define the order and grouping of tasks.
- Module: core building block; modules are invoked within tasks and are
  responisble for executing a specific action on the host;
- Inventory: file that defines groups of hosts that Ansible can target
- Variables: hosts variables, group variables, play variables; they can be used
  within playbook; they can be define in the Inventory, separate YAML files or
  in the Playbooks itself;
- Facts: Ansible gathers system info about hosts and stores it as 'facts' (OS,
  IP, hardware details) . They are collected automatically and can be used as
  varibales within playbooks.
- Handlers: specific tasks (for events like restarting a service) triggered by an event;
- Roles: way to organize and package playbooks, tasks, variables into reusable units;
- Templates: used for dynamically generated config files on the host;
- Vault: tool for encrypting sensitive data (passwords, secret keys) within Ansible playbook;

## Additional Resources

- https://www.vaultproject.io
- https://github.com/bitnami-labs/sealed-secrets
- https://learn.hashicorp.com/tutorials/vault/database-secrets?in=vault/db-credentials
- https://aws.amazon.com/secrets-manager/
- https://azure.microsoft.com/en-us/services/key-vault/
- https://docs.ansible.com/ansible/latest/getting_started/index.html
- https://docs.ansible.com/ansible/latest/cli/ansible-galaxy.html
