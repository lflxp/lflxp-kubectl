Kubectl is k8s cluster management module, which provides GUI interface for data query and resource operation

![s1.png](https://github.com/lflxp/lflxp-kubectl/blob/master/asset/s1.png)
![s2.png](https://github.com/lflxp/lflxp-kubectl/blob/master/asset/s2.png)

# Requirements

* kubernetes cluster
* config file (out cluster)
* k8s container (inner cluster)

# Install

```bash
git clone https://github.com/lflxp/lflxp-kubectl
cd lflxp-kubectl
make install
```

# Usage

> lflxp-kubectl

# Feature

- dashboard
- pod
- deployment
- service
- nodes

# KeyBindings

*  F1: Show keybinding help
*  F2: Dashboard View & back to Dashboard View && Refresh
*  F3: Pod View & back to Pod View && Refresh 
*  F4: Deployment View && Refresh
*  F5: Service View && Refresh
*  F6: Node View && Refresh
*  Space: search current view information 
*  Ctrl+C: Exit
*  Ctrl+L: Show log -tail 200 on F3 Pod View || msg refresh Log view
*  ↑ ↓: Move View
*  Enter: Commit Input/Quit current view(msg) 
*  Tab: Next View 

# Options

- describe
- delete
- search
- logs

# go mod problem

[go mod fix](https://segmentfault.com/a/1190000021077653?utm_source=tag-newest)