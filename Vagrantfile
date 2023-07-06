# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
    config.vm.box = "ubuntu/focal64"

    (1..3).each do |i|
        config.vm.define "k8s-node-#{i}" do |node|
            node.vm.network :private_network, ip: "192.168.60.#{i+10}"
            node.vm.hostname = "k8s-master-#{i}"
            node.vm.provider "virtualbox" do |vb|
                vb.memory = "4096"
            end
        end
    end

    ssh_pub_key = File.readlines("./ssh/id_rsa.pub").first.strip
    config.vm.provision 'shell', inline: 'mkdir -p /root/.ssh'
    config.vm.provision 'shell', inline: "echo #{ssh_pub_key} >> /root/.ssh/authorized_keys"
    config.vm.provision 'shell', inline: "echo #{ssh_pub_key} >> /home/vagrant/.ssh/authorized_keys", privileged: false
end