# Public Cloud

## Definition
It gives **programmable resources** that are available (loan). They are
**dynamic abilities** (can scale as you go) and you may pay as you go.

### Pros

- Trade capital expense for variable expense (no need to buy hardware)
  - pay as you consume resources
  - pay only for what you consume
- Benefit from massive economies of scale
  - benefit of massive data centers
- Stop guessing capacity (and develop faster)
  - scale up/down as required
  - no need to overprovision
- Increase speed and agility
  - experiment and develop faster
- Added value
  - focus on work that adds value
- Go global in minutes
  - multiple aws regions around the world
  - lower latence and a better experience for you customer
  - achieve high availability and disaster recovery

### Cons

- Where the client need/want on-premise solution
  - institutions, goverments, airlines
- Kind of expensive
  - what is total cost of ownership
- Not every business scales that much to benefit of autoscaling of the cloud

## Types

- Cloud
- Hybrid
- On-premises

### IaaS (Infrastructure as a Service)
- virtual machines on loan

### PaaS (Platform as a Service)
- VM, apache, tomcat, java are preinstalled

### SaaS (Service as a Service)
- gmail, image recognition

![cloud_services](./assets/cloud_services.png)

## AWS CAF
- Business Perspective (it finance, it strategy, benefits realization, busness risk management)
- People perspective (resource management, incentive management, career management, training management, organization change management)
- Governance Perspective (portoflio management, program and project management, business performance measurment, license management)
- Platform Perspective (compute, network, storage, database provisioning, systems and solution architecture, application development)
- Security Perspective (identity access manangement, detective control, infra security, data protection, incident response)
- Operation Perspective (service monitoring, app performance monitoring, resource inventory management, release management, reporting and analytics, disaster recovery, it service catalog)

Data center.<br>
Availability zone - multiple data centers in one area.<br>
Region - multiple AZ.<br>
AWS backbone network - local network between regions.<br>

S3 - object storage (you modify the whole object, not like the blob storage)
- WORM (write once, read many)
- avoid when:
  - long-term archival storage
  - block storage
  - frequently changes

Elastic Compute Cloud - EC2
- virtual machine
- AMI - amazon machine images

Virtual Private Cloud - VPC
- private network within AWS
- internet gateway
- nat gateway (located within public network)

Databases
- RDS
  - comes with engine on demand (psql, mysql, oracle, etc)
- Redshift
- Aurora
  - compatible with psql or mysql
  - better performance (3x better that psql ..?!)
- DinamoDB
  - manage table
- ElasticCache
- Neptune

Containers:
- Elastic container registry - ECR
- Elastic Container Service - ECS (managing containers)
- Elastic Kubernetes Service - EKS (managed k8s cluster)
- hosting on:
  - EC2
  - Fargate

Serverless - deploy code (go, ruby, ..); and wait event trigger.
- example:
  - on upload in s3 (s3 event)
  - execute labda function
    - process the file
  - save to (another) s3 bucket

Resources:
- https://wa.aws.amazon.com/wat.map.en.html
- https://wellarchitecuredlabs.com
- https://workshops.aws
- https://catalog.workshops.aws
