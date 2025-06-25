
# 🎮 Hope Engine

**Hope Engine** is a modular, lightweight 2D game engine built in Go using the Ebiten library. Designed for rapid prototyping and creative experimentation, it prioritizes simplicity, performance, and expandability — with plans for 3D features in future versions.

---

## ✨ Features (Phase 1)
- Custom rendering pipeline with Ebiten
- JSON-based tilemap system
- Modular core structure (rendering, input, entity management)
- Clean Go architecture for easy expansion
- Cross-platform support (Windows, Linux, macOS)

---

## 📁 Project Structure
```
HopeEngine/
├── assets/           // Tilemaps, sprites, audio
│   └── levels/       // Level files in JSON
├── engine/           // Game systems and modules
├── main.go           // Entry point and game loop
├── go.mod            // Go module setup
└── README.md         // You’re here!
```

---

## 🚀 Getting Started

1. Install Go: [https://go.dev/dl](https://go.dev/dl)  
2. Clone the repo:
   ```bash
   git clone https://github.com/hopeigbinosa123/hope-engine.git
   cd hope-engine
   ```
3. Run it:
   ```bash
   go run main.go
   ```

---

## 🧠 Next Goals
- Implement player movement and collisions
- Create an in-editor tile placer
- Add AI behaviors (👀 what if they remember you?)
- Build the save/export pipeline

---

## 💬 Author
Created with intent and vision by **Hope Igbinosa**  
*“Start simple. Scale when they believe.”*

