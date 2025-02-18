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

    // Fungsi untuk mendapatkan nilai token pada cookie
    function getCookie(name) {
        var cookieName = name + "=";
        var decodedCookie = decodeURIComponent(document.cookie);
        var cookieArray = decodedCookie.split(';');

        for (var i = 0; i < cookieArray.length; i++) {
            var cookie = cookieArray[i];
            while (cookie.charAt(0) === ' ') {
                cookie = cookie.substring(1);
            }
            if (cookie.indexOf(cookieName) === 0) {
                return cookie.substring(cookieName.length, cookie.length);
            }
        }
    }

    var token = getCookie('token');

    // Cek Token
    if (token) {
        // Token ada dalam cookie, lakukan tindakan yang sesuai
        // console.log('Token:', token);

        // Tembak Api check token untuk memastikan token masih aktif atau tidak
        var url = 'http://localhost:8080/api/v1/main-app/dashboard/check-token';
        fetch(url, {
            method: 'GET',
            headers: {
                'Authorization': 'Bearer ' + token
            }
        }).then(response => response.json()).then(hasil => {
            if(hasil.status == "Sukses")
            {
                window.location.href = "http://localhost:8080/halaman-aplikasi/dashboard";
            }
        }).catch(error => {
            console.log(error);
        });
    }

    jQuery(".btn-register").click(function() {
        // Redirect ke halaman pendaftaran
        window.location.href = "http://localhost:8080/";
    })

    jQuery('.btn-login').click(function() {
        // Get form data
        var email = jQuery(".log-email").val();
        var password = jQuery(".log-password").val();

        var formData = new FormData();
        formData.append('email', email);
        formData.append('password', password);

        // Kirim permintaan pembaruan produk ke API
        jQuery.ajax({
            url: 'http://localhost:8080/api/v1/autentikasi/login',
            type: 'POST',
            data: formData,
            processData: false,
            contentType: false,
            success: function(response) {
                if (response.status == 'Sukses')
                {
                    // Fungsi untuk mengatur cookie
                    function setCookie(name, value, days) {
                        var expires = "";

                        if (days) {
                            var date = new Date();
                            date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
                            expires = "; expires=" + date.toUTCString();
                        }

                        document.cookie = name + "=" + encodeURIComponent(value) + expires + "; path=/";
                    }

                    // Fungsi untuk mendapatkan nilai cookie
                    function getCookie(name) {
                        var cookieName = name + "=";
                        var decodedCookie = decodeURIComponent(document.cookie);
                        var cookieArray = decodedCookie.split(';');

                        for (var i = 0; i < cookieArray.length; i++) {
                            var cookie = cookieArray[i];
                            while (cookie.charAt(0) === ' ') {
                                cookie = cookie.substring(1);
                            }
                            if (cookie.indexOf(cookieName) === 0) {
                                return cookie.substring(cookieName.length, cookie.length);
                            }
                        }

                        return null;
                    }

                    // Simpan token ke cookie
                    var token = response.token;
                    setCookie('token', token, 14); // Simpan cookie selama 7 hari

                    // Tampilkan notifikasi toast sukses
                    jQuery('.pesan-sukses').text(response.message);
                    Toastify({
                        node: $("#success-notification-content")
                            .clone()
                            .removeClass("hidden")[0],
                        duration: 2000,
                        newWindow: true,
                        close: true,
                        gravity: "top",
                        position: "right",
                        stopOnFocus: true,
                    }).showToast();

                    // Redirect ke halaman dashboard admin
                    setTimeout(function() {
                        window.location.href = "http://localhost:8080/halaman-aplikasi/dashboard"; // Replace with your desired destination URL
                    }, 3000); // 3000 milliseconds = 3 seconds
                }else{
                    jQuery('.pesan-gagal').text(response.message);
                    Toastify({
                        node: $("#failed-notification-content")
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
            },
            error: function(xhr, status, error) {
                // Show the modal
                // Parsing responseText sebagai JSON
                var hpes = JSON.parse(xhr.responseText);
                jQuery('.pesan-gagal').text(hpes.message);
                Toastify({
                    node: $("#failed-notification-content")
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
        });
    })
})