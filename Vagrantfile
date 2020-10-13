Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/xenial64"

  config.vm.network "private_network", ip: "192.168.33.10"

  config.vm.synced_folder ".", "/vagrant/passnet-auth", id: "app"
  config.vm.synced_folder "../sagas/v1/saga-authentication", "/vagrant/saga", id: "saga"

  config.vm.provider "virtualbox" do |vb|
    vb.gui = true
    vb.memory = "2048"
  end

  config.vm.provision "set-env", type: "shell", inline: "export PATH=$PATH:/usr/local/go/bin", run: "always"
  config.vm.provision :docker
end
