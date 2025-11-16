const btnGenerar = document.getElementById('btnGenerar');
const btnVerTodas = document.getElementById('btnVerTodas');
const quoteContainer = document.getElementById('quoteContainer');
const quotesList = document.getElementById('quotesList');


async function generarCita() {
    try {
        const response = await fetch('/api/quote');
        if (!response.ok) {
            throw new Error('Error al generar la cita');
        }
        const data = await response.json();

        quotesList.style.display = 'none';

        quoteContainer.innerHTML = `
            <blockquote>
                ${data.quote}
            </blockquote>
            <p><small>Total de citas generadas: ${data.total}</small></p>
        `;
        quoteContainer.style.display = 'block';
    } catch (error) {
        console.error('Error:', error);
        quoteContainer.innerHTML = `
            <p style="color: red;">Error al generar la cita. Por favor, intenta de nuevo.</p>
        `;
    }
}

async function obtenerTodasLasCitas() {
    try {
        const response = await fetch('/api/quotes');
        if (!response.ok) {
            throw new Error('Error al obtener las citas');
        }
        const data = await response.json();

        quoteContainer.style.display = 'none';

        if (data.quotes && data.quotes.length > 0) {
            quotesList.innerHTML = `
                <h2>Todas las citas generadas (${data.total})</h2>
                <ul>
                    ${data.quotes.map(quote => `<li>${quote}</li>`).join('')}
                </ul>
            `;
        } else {
            quotesList.innerHTML = `
                <p>No hay citas generadas aún. ¡Genera tu primera cita!</p>
            `;
        }
        quotesList.style.display = 'block';
    } catch (error) {
        console.error('Error:', error);
        quotesList.innerHTML = `
            <p style="color: red;">Error al obtener las citas. Por favor, intenta de nuevo.</p>
        `;
    }
}

btnGenerar.addEventListener('click', generarCita);
btnVerTodas.addEventListener('click', obtenerTodasLasCitas);

