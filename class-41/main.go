/*
thread kivabe manage hoi thread ke exicute kore?
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
*/