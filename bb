Current Environment Details:
ATI has numerous applications which provide business services to both ATI users and partners. The applications selected for DR provide the most critical functions required to continue day to day operations. Required infrastructure, network design, storage, file shares, Databases, AD servers required for these core applications are present in the AWS Region, U.S. East(N.Virginia). This environment is planned to be made active and servers as production environment after migration. The components currently present in AWS environment are discussed below.

Network Services:
Network services provide capabilities to allow external and internal, systems and users to connect to application services. 
• Virtual Private Clouds (VPC) / Network Subnets – Base Networking has been staged and is ready to receive applications.
• Palo Alto Firewalls – Firewalls have been setup with rules necessary to support ingress and egress traffic.
• BGP routing – Routing has been established to allow both data replication and user access required for accessing applications.
Authentication services:
• Active Directory domain controllers – Domain controllers provide user authentication to the ATI applications. 3 Domain Controllers have been pre-staged in the AWS environment
• ADFS – ADFS provides SSO with 3rd party SAAS providers. 2 sets of ADFS servers (internal and proxy) have been pre-staged in the AWS environment, including proper networking rules and load balancers.

File Services:
• DFS Namespace servers – These servers control access to the ATI file servers by presenting access through share names which can coexist in production and DR. 2 Namespace servers are pre-staged for redundancy. DR shares are not accessible to end users until activated.
• DFS File Servers – These servers store and host the files accessed through the Namespace servers. 2 File servers are pre-staged for redundancy.
• AWS instance size and IP address reservations are pre-staged for the SFTP server.

Database Services:
• 2 SQL servers have been deployed to AWS, for the purposes of hosting the DB-ODSDW-PROD and DB-TOUCHSTONE-PROD instances. These servers are configured for SQL AlwaysON and are deployed across 2 availability zones to provide additional redundancy.
• 2 SQL servers have been deployed to AWS, for the purposes of hosting the DB-CORPODS-PROD and DB-ADHOC-PROD instances. These servers are configured for SQL AlwaysON and are deployed across 2 availability zones to provide additional redundancy.
• 2 SQL servers have been deployed to AWS, for the purposes of hosting the DB-DW-PROD instance. These servers are configured for SQL AlwaysON and are deployed across 2 availability zones to provide additional redundancy.
• SQL jobs have been pre-staged in the AWS servers and have been disabled.

Touchstone & InSync:
6 Touchstone and 6 InSync application webservers have been deployed to AWS from an ATI customized instance. ATI applications have been installed on these servers, along with all prerequisites required to perform all required services. These systems have been deployed across three availability zones to provide additional redundancy.
AWS Network Load balancers have been deployed and configured to provide request distribution to the Touchstone/InSync application servers.
A standalone web\application server is pre-staged in order to manage Scheduled tasks, required by the application.
An AWS RDS database is pre-staged to provide application state services.
AWS instance size and IP address reservations are pre-staged for the System Integration server (for Touchstone).
AWS instance size and IP address reservations are pre-staged for the Credentialing Database server (for InSync).

Non-clustered databases SSIS, SSRS, Alteryx, Tableau, Great Plains, BI360 are replicated using CloudEndure tool. These are available in AWS environment in near realtime.


===================================================================================================================================================
Recovery activities will include:
IT Service recovery
Network Services:
The infrastructure components Virtual Private Clouds (VPC) / Network Subnets, Palo Alto Firewalls fully configured, Security groups and BGP routing with required access for users, Loadbalancers, etc  are configured in AWS environment.
Recovery Process:
Changing External DNS entries supporting ADFS and SFTP.
Opening firewall access to allow inbound communication.

Authentication Services:
Authentication services provide capabilities to allow users to access applications and data. ATI currently provides these services through the use of Microsoft Active Directory and Active Directory Federation Services. 2 sets of ADFS servers(for SSO with 3rd party SAAS providers), 3 Domain controllers (for user authentication to the ATI applications) have been pre-staged in the AWS environment, including proper networking rules and load balancers. Data replication for these services use Microsoft native replication technologies. Changes made in production will be automatically replicated to AWS. 
Recovery Process:
Promoting the AWS domain controller to act as the Primary Domain Controller for the ATI environment, by seizing FSMO roles.
Opening external networking for the ADFS proxy servers and changing external DNS to point to the AWS external IP address for ADFS services.
File Services:
DFS shares and files are replicated to the DR site using Microsoft’s DFS-R function. SFTP server is replicated in its entirety using the CloudEndure replication utility. These servers in AWS are in sync to current production environment.
Recovery process:
Modifying DFS shares to allow Read\Write Access, making the AWS DFS environment act as a primary write space.
Executing the failover process for the SFTP server, in the CloudEndure portal
Opening external networking for the SFTP server and changing external DNS to point to the AWS external IP address for SFTP.

Application Recovery
Database services (Touchstone and Insync)
The production environment uses 3 SQL clusters, 2 SQL servers DB-ODSDW-PROD and DB-TOUCHSTONE-PROD, 2 SQL servers for B-CORPODS-PROD and DB-ADHOC-PROD, 2 SQL instances for DB-DW-PROD are configured with SQL AlwaysON and are deployed across 2 availability zones. SQL data is replicated to offline databases, using differential backups and log shipping. Backups and logs will be applied as received to meet recover point targets.
Applying the latest available transactions to the offline databases.
Bringing the databases online
Updating internal DNS entries, to redirect Database aliases to the AWS servers.
Initiating SQL Jobs.
Performing validation activities.

Touchstone & InSync:
Code changes to the Touchstone/InSync application servers are manually deployed to the AWS staged servers, following deployment to production servers in Netsource. Hence no further steps needed to get latest code for these applications.
The System Integration server and Credentialing Database server is replicated in its entirety using the CloudEndure replication utility. Changes made to the server, in production, are replicated to the DR site in near real time.
Recovery Process:
Executing the Cloud Endure failover process for the System Integration server(Touchstone) and Credentialing Database server(InSync).
Updating internal DNS entries, related to the Touchstone/InSync applications, to reference the load balancer pre-staged in AWS.
Starting application services on the Touchstone/InSync application servers
Starting all services and scheduled tasks on the Standalone web\application server
Performing validation activities.

SSIS, SSRS, Alteryx, Tableau, Great Plains, BI360, 
All the above database servers are replicated in their entirety using the CloudEndure replication utility. Changes made to these servers, in production, are replicated to the AWS environment in near real time.
Recovery Process:
Executing the Cloud Endure failover process for the database servers.
Updating internal DNS entries, as related to the DR IP address reservation.
=========================================================================================================================================
