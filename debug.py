import psutil
import subprocess

process = filter(lambda p: "easy.env" in p.name(), psutil.process_iter())
for i in process:
    cmd = ['dlv', '--listen=:2345', '--headless=true', '--api-version=2', '--check-go-version=false', '--only-same-user=false', 'attach', str(i.pid)]
    subprocess.run(cmd)