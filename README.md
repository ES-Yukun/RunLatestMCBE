# RunLatestMCBE
 
# Features
* ENGLISH
- This container image is image provider agnostic, updates to the latest version on reboot, and is easy to run.
* 日本語
- このコンテナイメージはイメージプロバイダーに依存せず、再起動時に最新バージョンにアップデートされ、簡単に実行できます。

# Usage

* minecraft folder: `/root/minecraft`
* buckup folder:    `/root/minecraft/buckup`

```bash
docker run --name minecraft-bedrock-server -p 19132:19132/udp docheio/minecraft-be:latest
```

# Author
 
* ES-Yukun
* Anylinks Japan
* yukun@team.anylinks.jp
 
# License
 
"RunLatestMCBE" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).