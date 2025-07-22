# TOTTI 2.0

![Screenshot 2025-04-20 235246](https://github.com/user-attachments/assets/4244022e-74fe-491a-9cd1-5f043711756f)

TOTTI is a blazing-fast and highly scalable flag submission system built specifically for attack-defense CTF competitions. Designed to handle **millions of flags per round**, TOTTI combines a performant Go backend with a sleek and responsive frontend built with Vite.js and React.

> ⚽ Named after Francesco Totti – a legend in attack. Because your team should attack like a legend too.

🚀 Features
-----------
- ⚡ Handles 1M+ flags/round – Extremely efficient and scalable
- 🧠 Smart UI – Visualize exploit and team performance via graphs and charts
- 🧪 Manual Submission – Easily test flags directly from the UI
- 📊 Realtime Stats – Track timelines, exploit performance, team results
- 🎨 Built with:
  - Backend: Go
  - Frontend: Vite.js, React, MUI, ECharts, X-Charts

🖥️ Frontend Overview
---------------------
- Custom Graphs:
  - 📈 TimelineGraph: Track submissions over ticks
  - 📊 AllTeamsGraph: Visualize each team's flag results
  - 🐞 ExploitGraph: See how each exploit performs
  - 🥧 TotalPieGraph: Global status distribution (Success, Pending, Expired, Error)

- Components:
  - Navbar: Filter by exploit and time range
  - ManualSubmission: Submit a flag manually via input field
  - CustomTitle: Branded header with the TOTTI logo

🧰 Backend Highlights
---------------------
- Written in Go for maximum throughput
- Optimized for concurrent processing
- Clean API to receive and handle flag batches or manual entries
- Designed to scale vertically and horizontally

🐳 Deployment (with Docker)
----------------------------
A simple Docker setup is included. To run:

```bash
docker build -t totti .
docker run -p 3000:3000 totti
```

Adjust port mapping and environment variables as needed.

