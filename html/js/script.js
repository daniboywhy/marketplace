document.addEventListener('DOMContentLoaded', function() {
    function atualizarCarrinho(inventory) {
        var tbody = document.querySelector('tbody');
        tbody.innerHTML = ''; // Limpa o conteúdo existente

        inventory.forEach(function(item) {
            var row = document.createElement('tr');

            var produtoTd = document.createElement('td');
            var produtoDiv = document.createElement('div');
            produtoDiv.classList.add('produto');

            var img = document.createElement('img');
            img.src = 'https://picsum.photos/100/120'; // URL da imagem
            img.alt = item.Nome;

            var infoDiv = document.createElement('div');
            infoDiv.classList.add('info');

            var nameDiv = document.createElement('div');
            nameDiv.classList.add('name');
            nameDiv.textContent = item.Nome;

            var categoryDiv = document.createElement('div');
            categoryDiv.classList.add('category');
            categoryDiv.textContent = 'Categoria'; // Ajuste conforme necessário

            infoDiv.appendChild(nameDiv);
            infoDiv.appendChild(categoryDiv);

            produtoDiv.appendChild(img);
            produtoDiv.appendChild(infoDiv);
            produtoTd.appendChild(produtoDiv);

            var precoTd = document.createElement('td');
            precoTd.textContent = 'R$ ' + item.Valor;

            var quantidadeTd = document.createElement('td');
            var qtyDiv = document.createElement('div');
            qtyDiv.classList.add('qty');

            var minusButton = document.createElement('button');
            minusButton.innerHTML = '<i class="bx bx-minus"></i>';

            var qtySpan = document.createElement('span');
            qtySpan.textContent = item.Quantity;

            var plusButton = document.createElement('button');
            plusButton.innerHTML = '<i class="bx bx-plus"></i>';

            qtyDiv.appendChild(minusButton);
            qtyDiv.appendChild(qtySpan);
            qtyDiv.appendChild(plusButton);
            quantidadeTd.appendChild(qtyDiv);

            var totalTd = document.createElement('td');
            var total = (parseFloat(item.Valor) * item.Quantity).toFixed(2);
            totalTd.textContent = 'R$ ' + total;

            var removeTd = document.createElement('td');
            var removeButton = document.createElement('button');
            removeButton.classList.add('remove');
            removeButton.innerHTML = '<i class="bx bx-x"></i>';

            removeTd.appendChild(removeButton);

            row.appendChild(produtoTd);
            row.appendChild(precoTd);
            row.appendChild(quantidadeTd);
            row.appendChild(totalTd);
            row.appendChild(removeTd);

            tbody.appendChild(row);
        });
    }

    function fetchInventory() {
        fetch('/inventory')
            .then(response => {
                if (!response.ok) {
                    throw new Error('Erro ao buscar inventário: ' + response.status);
                }
                return response.json();
            })
            .then(data => {
                atualizarCarrinho(data);
            })
            .catch(error => {
                console.error('Erro ao buscar inventário:', error);
            });
    }

    // Atualiza o carrinho ao carregar a página
    fetchInventory();

    // Você pode também adicionar um evento de clique para o botão "Atualizar Carrinho"
    var atualizarCarrinhoButton = document.querySelector('#btt_att_carrinho');
    if (atualizarCarrinhoButton) {
        atualizarCarrinhoButton.onclick = fetchInventory;
    }
});