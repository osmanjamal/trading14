-- إنشاء جدول الصفقات
CREATE TABLE IF NOT EXISTS trades (
    id SERIAL PRIMARY KEY,
    symbol VARCHAR(20) NOT NULL,
    action VARCHAR(10) NOT NULL,
    price DECIMAL(18, 8) NOT NULL,
    amount DECIMAL(18, 8) NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- إنشاء فهرس على عمود الرمز للبحث السريع
CREATE INDEX IF NOT EXISTS idx_trades_symbol ON trades(symbol);

-- إنشاء فهرس على عمود الوقت للترتيب والبحث السريع
CREATE INDEX IF NOT EXISTS idx_trades_timestamp ON trades(timestamp);

-- إنشاء جدول لحفظ لقطات الرصيد
CREATE TABLE IF NOT EXISTS balance_snapshots (
    id SERIAL PRIMARY KEY,
    currency VARCHAR(10) NOT NULL,
    amount DECIMAL(18, 8) NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- إنشاء فهرس على عمود العملة للبحث السريع
CREATE INDEX IF NOT EXISTS idx_balance_snapshots_currency ON balance_snapshots(currency);

-- إنشاء فهرس على عمود الوقت للترتيب والبحث السريع
CREATE INDEX IF NOT EXISTS idx_balance_snapshots_timestamp ON balance_snapshots(timestamp);
