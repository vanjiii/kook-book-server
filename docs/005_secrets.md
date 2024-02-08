# Secrets

## Challenges

- to store in vcs
- to manage across different environements
- authentication
- authorization
- manage application & user control
- to rotate
- for audit

`git ref-log` saves everything, and need to be cleaned if secrets leaked to the repo:<br>
- git-leaks
- git-secrets
- detect-secrets
- checkov

## Secrets Management Solutions

- Secrets at rest - how they are stored
- Secrets in transit - how they are transferred to the application (https for example)
- Secrets in used

![secret-manager](./assets/secret_manager.png)

**Tools**: <br>
- PGP encrypted files
- Ansible vault
  - sits next to your code
  - fast
  - no audit trail
  - no version control
- cloud native
  - aws secret manager (parameters store as a free alternative for aws)
    - data enc in rest
    - data enc in transit
    - rotation
    - audit
    - no cheap
  - azure key vault
  - k8s secrets
  - **top** Hashicorp Vault
    - use different backend (file, consul, etcd, db)
    - user identity
    - self managed cloud
    - has k8s operator
    - high availability (master-slave arch)
  - Github Secrets

## How to

1. App direct calls vault on **every** call or whatever (e.g. db conn establishement)
  - if vault is gone app fails
2. App is in Env, env calls vault
3. VaultSDK or vault agent (side car container)
  - has caching
  - sdk is calling the vault

Temporary credentials (dynamic vaults): JIT access

## Access Control

- CRUD
```
path "secret/dev {
  capabilities = ["create", "read", "update", "delete", "list"]
}
```

- Internal Vars (user can update its own password, nothing else)
```
path "auth/userpass/users/{{ identity.user.pas }} {
  capabilities = ["read"]
  allowed_parameters = {
    "password" = []
  }
}
```

- Wildcard match
```
path "secret/data/* {
  capabilities = ["create", "read", "update", "delete", "list"]
}
```

Reads:<br>
- https://learn.hashicorp.com/tutorials/vault/getting-started-first-secret?in=vault/getting-started
- https://learn.hashicorp.com/tutorials/vault/database-secrets?in=vault/db-credentials
