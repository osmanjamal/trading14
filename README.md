# نظام التداول الآلي

هذا المشروع هو نظام تداول آلي يعمل على مدار الساعة (24/7) ومصمم للعمل على خادم VPS. يستقبل إشارات من TradingView ويقوم بتنفيذ عمليات التداول تلقائيًا على منصة Binance.

## المتطلبات الأساسية

قبل البدء، تأكد من تثبيت البرامج والأدوات التالية:

1. [Go](https://golang.org/dl/) (الإصدار 1.17 أو أحدث)
2. [PostgreSQL](https://www.postgresql.org/download/) (الإصدار 12 أو أحدث)
3. [Git](https://git-scm.com/downloads)
4. [Docker](https://docs.docker.com/get-docker/) (اختياري، للنشر)

## إعداد المشروع

1. استنساخ المستودع:
   ```
   git clone https://github.com/osmanjamal/trading12
   cd trading-bot
   ```

2. تثبيت الاعتماديات:
   ```
   go mod download
   ```

3. إعداد قاعدة البيانات:
   - قم بإنشاء قاعدة بيانات جديدة في PostgreSQL
   - قم بتشغيل السكريبت الموجود في `database/schema.sql` لإنشاء الجداول اللازمة

4. إعداد ملف البيئة:
   قم بإنشاء ملف `.env` في الدليل الجذر للمشروع وأضف المتغيرات التالية:
   ```
   PORT=8080
   DATABASE_URL=postgres://username:password@localhost:5432/dbname
   EXCHANGE_API_KEY=your_binance_api_key
   EXCHANGE_SECRET_KEY=your_binance_secret_key
   LOG_LEVEL=info
   ```

## تشغيل المشروع

لتشغيل المشروع محليًا:

```
go run cmd/server/main.go
```

للبناء والتشغيل:

```
go build -o trading-bot cmd/server/main.go
./trading-bot
```

## النشر على VPS

1. قم بنقل الكود إلى VPS الخاص بك.
2. قم بتثبيت Go وPostgreSQL على VPS.
3. اتبع خطوات الإعداد والتشغيل المذكورة أعلاه.
4. استخدم أداة مثل `systemd` أو `supervisor` لضمان تشغيل البرنامج باستمرار.

مثال على ملف خدمة `systemd`:

```
[Unit]
Description=Trading Bot Service
After=network.target

[Service]
ExecStart=/path/to/trading-bot
WorkingDirectory=/path/to/project/directory
User=youruser
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
```

## هيكل المشروع

```
trading-bot/
│
├── cmd/
│   └── server/
│       └── main.go           # نقطة الدخول الرئيسية للتطبيق
│
├── internal/
│   ├── api/
│   │   ├── handlers.go       # معالجات HTTP
│   │   └── routes.go         # تعريف مسارات API
│   │
│   ├── bot/
│   │   ├── signal_processor.go  # معالجة إشارات التداول
│   │   └── trading_logic.go     # منطق التداول
│   │
│   ├── database/
│   │   ├── models.go         # نماذج قاعدة البيانات
│   │   └── operations.go     # عمليات قاعدة البيانات
│   │
│   ├── exchange/
│   │   ├── binance.go        # تنفيذ واجهة Binance
│   │   └── interface.go      # واجهة التبادل العامة
│   │
│   └── config/
│       └── config.go         # إدارة التكوين
│
├── pkg/
│   ├── logger/
│   │   └── logger.go         # وحدة التسجيل
│   │
│   └── utils/
│       └── helpers.go        # وظائف مساعدة
│
├── web/
│   ├── templates/
│   │   └── index.html        # قالب HTML الرئيسي
│   │
│   └── static/
│       ├── css/
│       │   └── main.css      # أنماط CSS
│       │
│       └── js/
│           └── app.js        # سكريبت JavaScript للواجهة الأمامية
│
├── tests/
│   ├── api_test.go           # اختبارات API
│   └── bot_test.go           # اختبارات منطق البوت
│
├── go.mod                    # تعريف وحدة Go ومتطلباتها
├── go.sum                    # قفل إصدارات الاعتماديات
├── Dockerfile                # تعريف صورة Docker
├── .gitignore                # قائمة الملفات المتجاهلة من Git
└── README.md                 # وثائق المشروع (هذا الملف)
```

## المكتبات والحزم المستخدمة

- [gorilla/mux](https://github.com/gorilla/mux): للتوجيه HTTP
- [lib/pq](https://github.com/lib/pq): برنامج تشغيل PostgreSQL
- [spf13/viper](https://github.com/spf13/viper): لإدارة التكوين
- [gorilla/websocket](https://github.com/gorilla/websocket): لدعم WebSocket
- [uber-go/zap](https://github.com/uber-go/zap): للتسجيل عالي الأداء

## الإسهام

نرحب بالمساهمات! يرجى قراءة `CONTRIBUTING.md` للحصول على التفاصيل حول عملية تقديم طلبات السحب.

## الترخيص

هذا المشروع مرخص بموجب رخصة MIT - انظر ملف `LICENSE` للحصول على التفاصيل.

## الدعم

إذا واجهت أي مشاكل أو كانت لديك أسئلة، يرجى فتح مشكلة في هذا المستودع.#   t r a d i n g 1 2  
 