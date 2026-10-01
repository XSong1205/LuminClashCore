# LuminClashCore

> **LuminClash 专用的 Mihomo (Clash.Meta) 跨平台内核自动化构建与同步仓库。**

---

## 🌟 特性 (Features)

- 🔄 **自动跟踪上游**：每 6 小时自动对齐 [MetaCubeX/mihomo](https://github.com/MetaCubeX/mihomo) 官方最新 Release。
- 📱 **Android 跨平台动态库**：
  - 采用 Android NDK 预编译 C-Shared 动态链接库 (`libclash.so`)；
  - 自动输出覆盖 `arm64-v8a`、`armeabi-v7a`、`x86_64` 的全套 `jniLibs.zip`；
  - 内置 `with_gvisor` 标签，原生支持 Android `VpnService` TUN 虚拟网卡转发。
- 💻 **Windows 桌面端二进制**：
  - 预编译输出 `windows-amd64` 与 `windows-arm64` 的 `mihomo.exe`。
- 📦 **自动化 Release 发布**：
  - 构建完成后自动发布与上游同名的 GitHub Release，附带 `sha256sums.txt` 完整性校验清单。

---

## 📥 产物集成指引 (Integration Guide)

### 1. Android 端集成 (LuminClash)
从最新 [Releases](../../releases) 页面下载 `LuminClashCore-android-jniLibs-<version>.zip`，解压后拷贝至工程目录：

```text
LuminClash/android/app/src/main/jniLibs/
├── arm64-v8a/
│   └── libclash.so
├── armeabi-v7a/
│   └── libclash.so
└── x86_64/
    └── libclash.so
```

在打包 Android APK (`flutter build apk`) 时，Gradle 将自动把对应 ABI 架构的 `libclash.so` 嵌入安装包。

### 2. Windows 桌面端集成
从 [Releases](../../releases) 页面下载 `LuminClashCore-windows-amd64-<version>.zip`，解压得到 `mihomo.exe`，放置于：
```text
LuminClash/assets/core/mihomo.exe
```

---

## 🛠 手动触发构建 (Manual Trigger)

在 GitHub 仓库页面进入 **Actions** -> 选择 **Sync Upstream and Build Core** -> 点击 **Run workflow**：
- 可以直接点击运行触发最新版本构建；
- 或在 `force_version` 中填写目标版本号（例如 `v1.19.32`）强制构建指定版本。

---

## 📄 开源许可 (License)

本项目遵循 [GPL-3.0 License](LICENSE)，与上游 Mihomo 保持一致。
