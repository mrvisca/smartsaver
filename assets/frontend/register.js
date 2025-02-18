// cek package jquery
jQuery( document ).ready(function() {
    // Menonaktifkan klik kanan
    jQuery(document).on("contextmenu", function(e) {
        e.preventDefault();
    });

    // Menonaktifkan F12 (DevTools)
    jQuery(document).keydown(function(e) {
        if (e.key === "F12" || (e.ctrlKey && e.shiftKey && e.key === "I") || (e.ctrlKey && e.key === "U")) {
            e.preventDefault();
        }
    });

    jQuery(".btn-login").click(function() {
        // Redirect ke halaman pendaftaran
        window.location.href = "http://localhost:8080/";
    })

    // Fungsi tombol register
    jQuery('.btn-daftar').click(function() {
        // Get form data
        var namelog = jQuery(".log-name").val();
        var email = jQuery(".log-email").val();
        var password = jQuery(".log-password").val();
        var pasconfirm = jQuery(".log-pasconfirm").val();
        var phone = jQuery(".log-phone").val();

        if (password === pasconfirm) {
            // Pembuatan form data untuk proses post
            var formData = new FormData();
            formData.append('name', namelog);
            formData.append('email', email);
            formData.append('password', password);
            formData.append('phone', phone);

            // Kirim permintaan pembaruan produk ke API
            jQuery.ajax({
                url: 'http://127.0.0.1:8080/api/v1/autentikasi/register',
                type: 'POST',
                data: formData,
                processData: false,
                contentType: false,
                success: function(response) {
                    // Show the modal
                    jQuery('.pesan-sukses').text(response.message);
                    Toastify({
                        node: $("#success-notification-content")
                            .clone()
                            .removeClass("hidden")[0],
                        duration: 3000,
                        newWindow: true,
                        close: true,
                        gravity: "top",
                        position: "right",
                        stopOnFocus: true,
                    }).showToast();

                    setTimeout(function() {
                        window.location.href = "http://localhost:8080/"; // Replace with your desired destination URL
                    }, 3000); // 3000 milliseconds = 3 seconds
                },
                error: function(xhr, status, message) {
                    // Show the modal
                    var hpes = JSON.parse(xhr.responseText);
                    jQuery('.pesan-gagal').text(hpes.message);
                    Toastify({
                        node: $("#failed-notification-content")
                            .clone()
                            .removeClass("hidden")[0],
                        duration: 4000,
                        newWindow: true,
                        close: true,
                        gravity: "top",
                        position: "right",
                        stopOnFocus: true,
                    }).showToast();

                    setTimeout(function() {
                        location.reload();
                    }, 8000); // 3000 milliseconds = 3 seconds
                }
            });
        }else{
            // Tampilkan Toast Konfirmasi Password
            Toastify({
                node: $("#konf-notification-content")
                    .clone()
                    .removeClass("hidden")[0],
                duration: 3000,
                newWindow: true,
                close: true,
                gravity: "top",
                position: "right",
                stopOnFocus: true,
            }).showToast();
        }
    })
})