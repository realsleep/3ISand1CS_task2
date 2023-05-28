# 3ISand1CS_task2

#### Prerequisites:
    - Install Docker on Ubuntu 22.04 Machine

### Ubuntu 22.04 environment:
   - ##### Step 1. Install and run Jenkins server
    - ###### docker run -p 8080:8080 -name jenkins -v jenkins_home:/var/jenkins_home  jenkins/jenkins
   - ##### Step 2. Set up Jenkins server (hint: http://[jenskins-server-ip]:8080). Create pipeline item. As a pipeline script you can use pipeline.groovy file.
   - ##### Step 3. Install plugins such as "CloudBees Disk Usage Simple", "Docker Plusgin", "Docker Pipeline", "Git plugin", "GitHub plugin" "Prometheus metrics" on Jenkins 
   - ##### Step 4. Set up trigger. Every time push event triggers Jenkins pipeline.
        - ###### Dashboard > Manage Jenkins > System > Github > Advanced and click on 'Specify another hook URL for GitHub configuration'. Copy given URL address. Go to GitHub repository 'Setting' tab and under 'Webhooks' create new Webhook and paste previous URL address
    - ##### Step 5. Set up Prometheus and Grafana server.
        - ###### On server create file prometheus.yml and paste content from repository prometheus.yml file
        - ###### docker run -d -p 9090:9090 -v /home/vagrant/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
        - ###### docker run -d -p 3000:3000 grafana/grafana-enterprise