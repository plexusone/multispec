# Author IRD

Guide the user through creating an Infrastructure Requirements Document that defines the deployment and operational environment.

## Overview

An IRD (Infrastructure Requirements Document) defines the infrastructure needed to deploy and operate the system. It covers compute, storage, networking, security, observability, and disaster recovery. It answers "Where and how will we run this?"

## Prerequisites

A TRD should exist to provide context on the technical architecture. The IRD translates technical requirements into infrastructure specifications.

## Workflow

### 1. Discovery Questions

Before starting the draft, gather essential information:

1. **Cloud Provider**: AWS, GCP, Azure, or multi-cloud?
2. **Environment Strategy**: How many environments? Dev, staging, prod?
3. **Scale Requirements**: Initial capacity and growth projections?
4. **Availability Target**: What uptime SLA is required?
5. **Compliance**: Any regulatory requirements (SOC2, HIPAA, GDPR)?
6. **Budget**: What are the cost constraints?

### 2. Initialize Draft

Use `start_draft` to create the IRD draft:

```
start_draft(project="<project-name>", spec_type="ird")
```

This returns the template and instructions.

### 3. Collaborative Authoring

Work with the user to fill in each section:

- **Infrastructure Overview**: Architecture diagram and environment strategy
- **Compute Requirements**: VMs, containers, serverless functions
- **Storage Requirements**: Databases, object storage, caching
- **Networking**: VPCs, subnets, load balancers, DNS, CDN
- **Security**: IAM, secrets management, encryption, compliance
- **Observability**: Logging, metrics, tracing, alerting
- **Availability & DR**: Multi-region strategy, failover, backup
- **Capacity Planning**: Initial sizing and scaling strategy
- **CI/CD Infrastructure**: Pipeline, artifact storage, deployment

### 4. Save Progress

After each major section, save the draft:

```
update_draft(project="<project-name>", spec_type="ird", content="<full-content>")
```

### 5. Evaluate Quality

When the draft feels complete, run evaluation:

```
eval_draft(project="<project-name>", spec_type="ird")
```

Review the findings and score. Address any critical or high severity issues.

### 6. Iterate

Based on evaluation feedback:

1. Ensure infrastructure diagram covers all components
2. Verify security controls include IAM, encryption, and secrets
3. Check availability targets with RTO/RPO defined
4. Add observability stack with dashboards and alerts
5. Include capacity planning with cost estimates
6. Document IaC approach and runbooks
7. Re-evaluate until passing

### 7. Finalize

Once evaluation passes:

```
finalize_draft(project="<project-name>", spec_type="ird")
```

This promotes the draft to the final `technical/ird.md` location.

## Evaluation Criteria

The IRD is evaluated on:

| Category | Weight | What We Look For |
|----------|--------|------------------|
| Architecture Completeness | 20% | All components and connections documented |
| Security Design | 20% | IAM, encryption, secrets, and compliance |
| Availability & DR | 15% | SLAs, multi-region strategy, DR procedures |
| Observability | 15% | Logging, metrics, tracing, and alerting |
| Capacity & Cost | 15% | Projections with cost breakdown |
| Operability | 15% | IaC approach, CI/CD, and runbooks |

A passing score is 7.0+ with no critical or high severity findings.

## Tips

- Use tables for compute, storage, and networking specifications
- Include cost estimates for each resource category
- Security should cover both at-rest and in-transit encryption
- Define specific SLAs (uptime, RTO, RPO)
- Document scaling triggers and thresholds
- Include runbook references for common operations
- Consider IaC tooling (Terraform, Pulumi, CloudFormation)

## Infrastructure Diagram Format

```
┌─────────────────────────────────────────────────────────────────────┐
│                        Cloud Provider / Region                       │
├─────────────────────────────────────────────────────────────────────┤
│                                                                      │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐          │
│  │   Frontend   │    │    API       │    │   Database   │          │
│  │   (CDN)      │───▶│   (Compute)  │───▶│   (Storage)  │          │
│  └──────────────┘    └──────────────┘    └──────────────┘          │
│                                                                      │
└─────────────────────────────────────────────────────────────────────┘
```

## Resource Table Format

| Component | Type | Size | Count | Scaling |
|-----------|------|------|-------|---------|
| API Server | Container | 2 vCPU, 4GB | 3 | Auto (2-10) |
| Database | RDS PostgreSQL | db.r5.large | 1 Primary + 1 Replica | Manual |
| Cache | ElastiCache Redis | cache.r5.large | 2 nodes | Manual |

## Cost Estimation Format

| Resource | Monthly Cost | Notes |
|----------|--------------|-------|
| Compute | $500 | 3x API servers |
| Database | $400 | RDS with replica |
| Storage | $100 | S3 + EBS |
| Network | $50 | NAT Gateway + data transfer |
| **Total** | **$1,050** | |

## Availability Targets

| Metric | Target |
|--------|--------|
| Uptime | 99.9% |
| RTO | 1 hour |
| RPO | 15 minutes |

## Next Steps

After IRD completion:

1. **Infrastructure Provisioning**: Set up environments using IaC
2. **CI/CD Setup**: Configure deployment pipelines
3. **Monitoring Setup**: Deploy observability stack
