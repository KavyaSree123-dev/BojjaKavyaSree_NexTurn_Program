// script.js
//  JavaScript

document.addEventListener("DOMContentLoaded", () => {
    const expenseForm = document.getElementById("expense-form");
    const expenseTableBody = document.getElementById("expense-table-body");
    const categoryDetails = document.getElementById("category-details");
    const expenseGraph = document.getElementById("expense-graph").getContext("2d");

    let expenses = JSON.parse(localStorage.getItem("expenses")) || [];

    const updateUI = () => {
        expenseTableBody.innerHTML = "";
        const categoryTotals = {};

        expenses.forEach((expense, index) => {
            const row = document.createElement("tr");

            row.innerHTML = `
                <td>${expense.amount}</td>
                <td>${expense.note}</td>
                <td>${expense.group}</td>
                <td><button onclick="deleteExpense(${index})">Remove</button></td>
            `;

            expenseTableBody.appendChild(row);

            categoryTotals[expense.group] = (categoryTotals[expense.group] || 0) + parseFloat(expense.amount);
        });

        categoryDetails.innerHTML = "";
        for (let group in categoryTotals) {
            const li = document.createElement("li");
            li.textContent = `${group}: $${categoryTotals[group].toFixed(2)}`;
            categoryDetails.appendChild(li);
        }

        renderGraph(categoryTotals);
    };

    const renderGraph = (data) => {
        const labels = Object.keys(data);
        const values = Object.values(data);

        new Chart(expenseGraph, {
            type: "pie",
            data: {
                labels,
                datasets: [
                    {
                        label: "Expenditure by Group",
                        data: values,
                        backgroundColor: ["#ff6384", "#36a2eb", "#cc65fe"],
                    },
                ],
            },
        });
    };

    const deleteExpense = (index) => {
        expenses.splice(index, 1);
        localStorage.setItem("expenses", JSON.stringify(expenses));
        updateUI();
    };

    expenseForm.addEventListener("submit", (e) => {
        e.preventDefault();

        const amount = document.getElementById("amount").value;
        const note = document.getElementById("note").value;
        const group = document.getElementById("group").value;

        expenses.push({ amount, note, group });
        localStorage.setItem("expenses", JSON.stringify(expenses));
        updateUI();

        expenseForm.reset();
    });

    window.deleteExpense = deleteExpense;

    updateUI();
});
