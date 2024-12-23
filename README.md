# CS9549

Final state of our Platform-as-a-Service (PaaS) system built on Kubernetes. The PaaS enables users to manage namespaces, deploy applications, scale resources, and remove deployments using an intuitive interface and robust backend architecture. This system integrates Kubernetes features with custom services to provide automation, isolation, and scalability for user applications (service providers).
Here the Service Discovery/Registry will be the kubernetes control plane. We created a custom user management microservice and deployment management microservice to provide apis for the UI side to enable customers(service providers) to deploy their applications, discover how many of their deployments are working, manage their deployments and get a endpoint from which they can interact with their custom applications that they deployed via our UI.

We are building something similar to qovery.com

Overall System Architecture
The entire application is deployed on kubernetes orchestration. Major components include:
● User Management Service: Handles user authentication and data.
● Deployment Management Service: Automates application deployment and scaling.
● Customer Services: Interfaces with Kubernetes to manage namespaces and resources.

![architecture](images\overall-arch.png)

Functional Requirements
User Management Service:
1. User Registration: Allows new users to sign up by providing username, email, and password.
2. Authentication: Secure login/logout with JWT.
3. User Account Management: Includes profile updates, password reset mechanisms, and account recovery.
4. API for Integration: Endpoints for creating, updating, and retrieving user data.
Deployment Management Service:
1. Automated Deployment: Automated deployment to environments (staging/production).
2. CI/CD and Version Control: Pulls latest changes from Git repositories and deploys automatically on single click from UI - Continuous deployment also possible without users intervention with just minor extra codes.
3. Rollback Capability: Supports rolling back to previous application versions.
4. Resource Provisioning: Allocates Kubernetes resources as required.
Customer Services:
1. Namespace Isolation: Assigns unique namespaces to each customer.
2. Application Deployment: Deploys applications via UI or API using container images.
3. Resource Monitoring: Provides real-time CPU, memory, and storage usage statistics.
4. Scaling Applications: Scales deployments based on customer requirements.
5. Git Integration: Enables automated deployments triggered by repository commits.
Quality Requirements
1. Security: Implements RBAC, encryption, and namespace isolation.
2. Usability: User-friendly UI for deploying, scaling, and managing applications.
3. Reliability: Ensures high uptime with robust error handling.
4. Scalability: Supports growing user and application demands.
5. Auditability: Logs user actions and deployments for troubleshooting.

Architecturally Significant Requirements (ASRs)
● Namespace Isolation: Ensures user-specific resource management and security.
● Scalability: Supports dynamic allocation of resources for customer applications.
● Rollback and CI/CD: Facilitates seamless version control and deployment pipelines.

Detailed Requirements are defined here: 
[CS 9549 Project Spreadsheet.xlsx](https://uwoca-my.sharepoint.com/:x:/g/personal/cezeagwu_uwo_ca/EUYlloT5zYJCuiRpmNhY8zcBCqgXU0hsw_8zYPbfAra9hA?e=YakXiA)

## Walkthrough

Login Page
![login](images\login-page.png)

Dashboard
![dashboard](images\dashboard.png)

Repo Scout
![repo-scout](images\repo-scout.png)

Create Deployments
![create-deployments](images\create-deployment.png)

Deployments
![deployments](images\deployments-page.png)

Deployment Details
![deployment-details](images\deployment-details.png)

Update Replicas
![update-replicas](images\update-replicas.png)

