# ZKeenIP Rulesets

Списки zkeenip от [JamesZero](https://github.com/jameszeroX/zkeen-ip) в формате mrs для использования с ядром Mihomo.

Пример использования:


```YAML
anchors:
  a1: &ipcidr { type: http, format: mrs, behavior: ipcidr, interval: 86400 }

rule-providers:
  akamai@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/akamai@ipcidr.mrs }
  amazon@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/amazon@ipcidr.mrs }
  arelion@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/arelion@ipcidr.mrs }
  azure@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/azure@ipcidr.mrs }
  bunnycdn@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/bunnycdn@ipcidr.mrs }
  cdn77@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/cdn77@ipcidr.mrs }
  cloudflare@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/cloudflare@ipcidr.mrs }
  colocrossing@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/colocrossing@ipcidr.mrs }
  contabo@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/contabo@ipcidr.mrs }
  digitalocean@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/digitalocean@ipcidr.mrs }
  fastly@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/fastly@ipcidr.mrs }
  frantech@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/frantech@ipcidr.mrs }
  gcore@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/gcore@ipcidr.mrs }
  google@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/google@ipcidr.mrs }
  hetzner@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/hetzner@ipcidr.mrs }
  leaseweb@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/leaseweb@ipcidr.mrs }
  linode@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/linode@ipcidr.mrs }
  liquidweb@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/liquidweb@ipcidr.mrs }
  mega@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/mega@ipcidr.mrs }
  meta@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/meta@ipcidr.mrs }
  oracle@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/oracle@ipcidr.mrs }
  ovh@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/ovh@ipcidr.mrs }
  scaleway@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/scaleway@ipcidr.mrs }
  telegram@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/telegram@ipcidr.mrs }
  vultr@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/vultr@ipcidr.mrs }
  youtube@ipcidr: { <<: *ipcidr, url: https://github.com/zxc-rv/zkeenip-rulesets/releases/latest/download/youtube@ipcidr.mrs }

rules:
  - RULE-SET,akamai@ipcidr,PROXY
  - RULE-SET,amazon@ipcidr,PROXY
  - RULE-SET,arelion@ipcidr,PROXY
  - RULE-SET,azure@ipcidr,PROXY
  - RULE-SET,bunnycdn@ipcidr,PROXY
  - RULE-SET,cdn77@ipcidr,PROXY
  - RULE-SET,cloudflare@ipcidr,PROXY
  - RULE-SET,colocrossing@ipcidr,PROXY
  - RULE-SET,contabo@ipcidr,PROXY
  - RULE-SET,digitalocean@ipcidr,PROXY
  - RULE-SET,fastly@ipcidr,PROXY
  - RULE-SET,frantech@ipcidr,PROXY
  - RULE-SET,gcore@ipcidr,PROXY
  - RULE-SET,google@ipcidr,PROXY
  - RULE-SET,hetzner@ipcidr,PROXY
  - RULE-SET,leaseweb@ipcidr,PROXY
  - RULE-SET,linode@ipcidr,PROXY
  - RULE-SET,liquidweb@ipcidr,PROXY
  - RULE-SET,mega@ipcidr,PROXY
  - RULE-SET,meta@ipcidr,PROXY
  - RULE-SET,oracle@ipcidr,PROXY
  - RULE-SET,ovh@ipcidr,PROXY
  - RULE-SET,scaleway@ipcidr,PROXY
  - RULE-SET,telegram@ipcidr,PROXY
  - RULE-SET,vultr@ipcidr,PROXY
  - RULE-SET,youtube@ipcidr,PROXY
  - MATCH,DIRECT
