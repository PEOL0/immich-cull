# 📸 Immich Cull

A simple Go application that connects to your [Immich](https://immich.app/) instance to download favourite assets from albums. Specifically created because I feel it is slow to sift through images in [RawTherapee](https://rawtherapee.com/).

## ✨ Features

- 🎯 **Album Selection**: Interactive album browsing and selection
- ⭐ **Favourites Filter**: Automatically identifies and downloads only favorited assets
- 🚀 **Concurrent Downloads**: Multi-threaded downloading with rate limiting to respect server resources
- 🔄 **Retry Logic**: Error handling with automatic retry for failed downloads
- 🛡️ **Environment Configuration**: Secure API key management via `.env` files

## 🚀 Getting Started

### Prerequisites

- Go 1.19 or higher
- An active Immich instance
- Immich API key

### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/PEOL0/immich-cull.git
   cd immich-cull
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Configure environment**

   Create a `.env` file in the project root:

   ```bash
   cp .env.example .env
   ```

   Edit `.env` with your Immich details:

   ```properties
   ImmichKey=your_immich_api_key_here
   ImmichURL=http://your-immich-instance
   ```

### Building

Build the application:

```bash
go build -o immich-cull src/main.go
```

## 🎮 Usage

Run the application with a target directory:

```bash
./immich-cull -D /path/to/download/directory
```

### Workflow

1. **Directory Selection**: The app will prompt you to confirm or modify the download directory
2. **Album Selection**: Browse through your available albums and select one by entering its number
3. **Automatic Processing**: The app will:
   - Scan the selected album for favourite assets
   - Display the count of favourites found
   - Download all favourite assets to your specified directory

### Command Line Options

- `-D`: Specify the download directory path

## 🔧 Configuration

### Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `ImmichKey` | Your Immich API key | `your_api_key_here` |
| `ImmichURL` | Your Immich instance URL | `http://your-immich-instance` |

### Performance Tuning

The application includes several performance optimizations:

- **Concurrent Downloads**: Limited to 3 simultaneous downloads to prevent server overload
- **Retry Logic**: Up to 3 retry attempts for failed downloads
- **Rate Limiting**: Built-in delays between requests
- **Timeout Handling**: 60-second timeout for download requests

## 🛠️ Development

### Dependencies

- `github.com/joho/godotenv` - Environment variable management
