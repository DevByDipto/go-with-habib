/*
1. thread kivabe manage hoi thread ke exicute kore?
✅ CPU → OS চালায়
   (CPU, OS kernel এর code execute করে)

✅ OS (Kernel) → Thread manage করে
   - কয়টা thread হবে
   - কোন thread কখন চলবে
   - thread এর priority কত
   - thread কে কত CPU time পাবে
   etc. এসব decision নেয়

✅ CPU → Thread execute করে (OS এর instruction অনুযায়ী)
   (Kernel যে thread select করে, CPU সেটা execute করে)

2. akta process er moddhe j akadhik thake ,kon thread kake exicute kortese ,koita thread ase aigulo ke hisab rakhe ?
ans: akta process er main thread stack heap code segment data segment vartual cpu etc er khobor process nije rakhe but akta process er moddhe main thread bade ar koita thread ase tar khobor rakhe sudu os er karnel

*/