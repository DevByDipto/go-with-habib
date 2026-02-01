/*
1. Hardware kii ?
কম্পিউটারের RAM, HDD, CPU—এই সবকিছুকে একত্রে এবং আলাদাভাবেও Hardware (হার্ডওয়্যার) বলা হয়।

সহজ কথায় বলতে গেলে: কম্পিউটারের যে অংশগুলো আপনি ছুঁতে পারেন (Physical parts) এবং যা শক্ত বা কঠিন পদার্থ দিয়ে তৈরি, সেগুলোই হার্ডওয়্যার।
2. process er moddhe context switching kivabe hoi ?
3. os threde vs process thred kii ?
4.core kii ?
5. jodi akadhik vartual cpu thake akadhik processor o tate akadhik thred thake to os full jinista ke manage kore kivabe
6. thred thake koi ? process thred er vitor thake naki thred processer vitor ?
7.thread kii ?
ans: thread holo light weight process jaa process er cpu,ram,HDD shob kichur sahajjche process e ram er moddhe thaka code segment er code er small portion exicute kore.


Thread হলো light weight process.এটি প্রসেসের ভেতরেই থাকে এবং প্রসেসের চেয়ে অনেক কম মেমোরি খরচ করে.Process-এর RAM এবং CPU-এর সাহায্যে কাজ করে.থ্রেড প্রসেসের জন্য বরাদ্দকৃত RAM-এর মেমোরি (Code, Data, Heap) শেয়ার করে এবং CPU-তে গিয়ে রান হয়.Code Segment-এর কোডের Small Portion এক্সিকিউট করে.বড় একটি প্রোগ্রামের নির্দিষ্ট কোনো একটি কাজ (যেমন শুধু ডাউনলোড করা বা শুধু টাইপিং হ্যান্ডেল করা) একটি থ্রেড সম্পন্ন করে।
*/