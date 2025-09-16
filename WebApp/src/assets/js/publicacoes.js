$('#nova-publicacao').on('submit', criarPublicacao);

$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao);

$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);


async function criarPublicacao(evento) {
    evento.preventDefault();

    const titulo = $('#titulo').val();
    const conteudo = $('#conteudo').val();

    const dados = new URLSearchParams();
    dados.append('titulo', titulo);
    dados.append('conteudo', conteudo);

    try {
        await fetch('/publicacoes', {
            method: 'POST',
            body: dados, // fetch já coloca o Content-Type correto automaticamente
        });

        window.location.href = "/home";

    } catch (erro) {
        Swal.fire("Erro ao criar publicação. Tente novamente mais tarde.");
    }
}

function curtirPublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');


    elementoClicado.prop('disabled', true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST",
    }).done(function () {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());
        contadorDeCurtidas.text(quantidadeDeCurtidas + 1);

        elementoClicado.addClass('descurtir-publicacao');
        elementoClicado.addClass('text-danger');
        elementoClicado.removeClass('curtir-publicacao');

    }).fail(function () {
        Swal.fire("Erro ao curtir a publicação. Tente novamente mais tarde.");
    }).always(function () {
        elementoClicado.prop('disabled', false);
    });
}

function descurtirPublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');


    elementoClicado.prop('disabled', true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST",
    }).done(function () {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());
        contadorDeCurtidas.text(quantidadeDeCurtidas - 1);

        elementoClicado.removeClass('descurtir-publicacao');
        elementoClicado.removeClass('text-danger');
        elementoClicado.addClass('curtir-publicacao');

    }).fail(function () {
        Swal.fire("Erro ao descurtir a publicação. Tente novamente mais tarde.");
    }).always(function () {
        elementoClicado.prop('disabled', false);
    });
}

function atualizarPublicacao(evento) {
    $(this).prop('disabled', true);

    const publicacaoId = $(this).data('publicacao-id');

    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        },
    }).done(function () {
        Swal.fire({
            title: 'Publicação atualizada!',
            text: 'Sua publicação foi atualizada com sucesso.',
            icon: 'success',
            confirmButtonText: 'OK'
        }).then(() => {
            window.location.href = "/home";
        });
    }).fail(function () {
        Swal.fire("Erro ao atualizar publicação. Tente novamente mais tarde.");
    }).always(function () {
        $('#atualizar-publicacao').prop('disabled', false);
    });

}

function deletarPublicacao(evento) {
    evento.preventDefault();

    Swal.fire({
        title: 'Tem certeza?',
        text: "Você não poderá recuperar essa publicação!",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#d33',
        cancelButtonColor: '#6b21a8',
        confirmButtonText: 'Sim, deletar!'
    }).then(function (confirmacao) {
        if (!confirmacao.value) return;

        const elementoClicado = $(evento.target);
        const publicacao = elementoClicado.closest('div');
        const publicacaoId = publicacao.data('publicacao-id');

        $.ajax({
            url: `/publicacoes/${publicacaoId}`,
            method: "DELETE",
        }).done(function () {
            publicacao.fadeOut("slow", function () {
                $(this).remove();
            });
        }).fail(function () {
            Swal.fire("Erro ao deletar a publicação. Tente novamente mais tarde.");
        });

    });
}