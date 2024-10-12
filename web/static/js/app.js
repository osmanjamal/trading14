document.addEventListener('DOMContentLoaded', function() {
    fetchTrades();
    fetchBalance();
    setInterval(fetchTrades, 60000); // Refresh every minute
    setInterval(fetchBalance, 60000);
});

function fetchTrades() {
    fetch('/api/trades')
        .then(response => response.json())
        .then(trades => {
            const tradesDiv = document.getElementById('trades');
            tradesDiv.innerHTML = '<h2>Recent Trades</h2>';
            const table = document.createElement('table');
            table.innerHTML = `
                <tr>
                    <th>Symbol</th>
                    <th>Action</th>
                    <th>Price</th>
                    <th>Amount</th>
                    <th>Timestamp</th>
                </tr>
            `;
            trades.forEach(trade => {
                const row = table.insertRow();
                row.innerHTML = `
                    <td>${trade.symbol}</td>
                    <td>${trade.action}</td>
                    <td>${trade.price}</td>
                    <td>${trade.amount}</td>
                    <td>${new Date(trade.timestamp).toLocaleString()}</td>
                `;
            });
            tradesDiv.appendChild(table);
        })
        .catch(error => console.error('Error fetching trades:', error));
}

function fetchBalance() {
    fetch('/api/balance')
        .then(response => response.json())
        .then(balance => {
            const balanceDiv = document.getElementById('balance');
            balanceDiv.innerHTML = '<h2>Current Balance</h2>';
            const ul = document.createElement('ul');
            for (const [currency, amount] of Object.entries(balance)) {
                const li = document.createElement('li');
                li.textContent = `${currency}: ${amount}`;
                ul.appendChild(li);
            }
            balanceDiv.appendChild(ul);
        })
        .catch(error => console.error('Error fetching balance:', error));
}