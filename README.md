# 3ISand1CS_task2

#### Prerequisites:
    - Install Docker on Ubuntu 22.04 Machine

### Ubuntu 22.04 environment:
   - ##### Step 1. Install and run Jenkins server
```console
curl -fsSL https://pkg.jenkins.io/debian-stable/jenkins.io-2023.key | sudo tee \
  /usr/share/keyrings/jenkins-keyring.asc > /dev/null
echo deb [signed-by=/usr/share/keyrings/jenkins-keyring.asc] \
  https://pkg.jenkins.io/debian-stable binary/ | sudo tee \
  /etc/apt/sources.list.d/jenkins.list > /dev/null
sudo apt-get update
sudo apt-get install jenkins
sudo apt update
sudo apt install openjdk-11-jre
java -version
sudo systemctl enable jenkins
sudo systemctl start jenkins
sudo systemctl status jenkins
```
   - ##### Step 2. Set up Jenkins server (URL: ``http://[jenskins-server-ip]:8080``). Hint: Don't forget to open port, you can use command ``ufw allow 8080``). Create new user. Create pipeline item. As a pipeline script you can use ``pipeline.groovy`` file.
   - ##### Step 3. Install plugins such as "Go", "Docker plugin"  on Jenkins.
   - ##### Step 4. Give Jenkins acceses to Docker 
``` console
sudo usermod -aG docker jenkins
sudo service jenkins restart
```
   - ##### Step 5. Set up trigger. Every time push event triggers Jenkins pipeline.
        - ###### Dashboard > Manage Jenkins > System > Github > Advanced and click on 'Specify another hook URL for GitHub configuration'. Copy given URL address. Go to GitHub repository 'Setting' tab and under 'Webhooks' create new Webhook and paste previous URL address
    - ##### Step 7. Set up Prometheus and Grafana server.
        - ###### On server create file prometheus.yml and paste content from repository prometheus.yml file
        - ###### docker run -d -p 9090:9090 -v ~/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
        - ###### docker run -d -p 3000:3000 grafana/grafana-enterprise