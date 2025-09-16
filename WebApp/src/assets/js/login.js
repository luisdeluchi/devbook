$('#login').on('submit', fazerLogin);

async function fazerLogin(evento) {
    evento.preventDefault();
    const email = $('#email').val();
    const senha = $('#senha').val();
    console.log('Email:', email);
    console.log('Senha:', senha); // Veja se não está vazio!

    try {
        const resposta = await fetch("/login", {
            method: "POST",
            headers: { "Content-Type": "application/x-www-form-urlencoded" },
            body: new URLSearchParams({ email, senha })
        });

        if (!resposta.ok) {
            if (resposta.status === 401) {
                await Swal.fire({
                    title: 'Falha no login',
                    text: 'E-mail ou senha incorretos.',
                    icon: 'error',
                    confirmButtonText: 'OK'
                });
                return;
            }
            throw new Error(`Erro inesperado: ${resposta.status}`);
        }

        await Swal.fire({
            title: 'Login realizado!',
            text: 'Você será redirecionado para a página inicial.',
            icon: 'success',
            confirmButtonText: 'OK'
        });

        window.location.href = "/home";

    } catch (erro) {
        console.error('Erro no login:', erro);
        await Swal.fire({
            title: 'Erro',
            text: 'Erro ao fazer login! Tente novamente.',
            icon: 'error',
            confirmButtonText: 'OK'
        });
    }
}
