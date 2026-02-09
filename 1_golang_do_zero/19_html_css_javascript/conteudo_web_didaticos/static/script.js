// Aguardar carregamento completo do DOM
document.addEventListener('DOMContentLoaded', function() {
    console.log('Página carregada e pronta!');
    
    // Selecionar elementos
    const botaoAlterar = document.getElementById('botao-alterar');
    const mensagem = document.getElementById('mensagem');
    const tituloPrincipal = document.getElementById('titulo-principal');
    
    // Adicionar evento de clique ao botão
    botaoAlterar.addEventListener('click', function() {
        // Alterar texto da mensagem
        mensagem.textContent = 'Mensagem alterada pelo JavaScript!';
        
        // Alterar cor da mensagem
        mensagem.style.color = '#4CAF50';
        mensagem.style.borderLeftColor = '#4CAF50';
        mensagem.style.backgroundColor = '#e8f5e9';
        
        // Alterar título
        tituloPrincipal.textContent = 'Título Alterado!';
        tituloPrincipal.style.color = '#FF9800';
        
        // Desabilitar botão após primeiro clique
        botaoAlterar.disabled = true;
        botaoAlterar.textContent = 'Já Clicado!';
        botaoAlterar.style.backgroundColor = '#999';
    });
    
    // Exemplo de manipulação adicional
    console.log('Elementos encontrados:', {
        botao: botaoAlterar !== null,
        mensagem: mensagem !== null,
        titulo: tituloPrincipal !== null
    });
});
