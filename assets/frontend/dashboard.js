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

        console.log(token)

        // Tembak Api check token untuk memastikan token masih aktif atau tidak
        var url = 'http://localhost:8080/api/v1/main-app/dashboard/check-token';
        fetch(url, {
            method: 'GET',
            headers: {
                'Authorization': 'Bearer ' + token
            }
        }).then(response => response.json()).then(hasil => {
            jQuery('.nama-user').text(hasil.data.name)
            jQuery('.role-user').text(hasil.data.role_name)
        }).catch(error => {
            console.log(error);
            window.location.href = "http://localhost:8080/";
        });

        jQuery('.btn-logout').click(function() {
            var formData = new FormData();
            
            jQuery.ajax({
                url: 'http://localhost:8080/api/v1/main-app/logout', // Sesuaikan dengan URL endpoint logout di server Anda
                type: 'POST',
                data: formData,
                processData: false,
                contentType: false,
                success: function(response) {
                    if (response.status == 'Sukses') {
                        // Fungsi untuk menghapus cookie
                        function deleteCookie(name) {
                            document.cookie = name + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
                        }
                        
                        // Hapus token dari cookie dan localStorage
                        deleteCookie('token');
                        localStorage.removeItem('token');
                        
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
                        
                        // Redirect ke halaman login
                        setTimeout(function() {
                            window.location.href = "http://localhost:8080/";
                        }, 3000);
                    }
                },
                error: function(xhr, status, error) {
                    var hpes = JSON.parse(xhr.responseText);
                    jQuery('.pesan-gagal').text(hpes.message);
                    Toastify({
                        node: $("#failed-notification-content")
                            .clone()
                            .removeClass("hidden")[0],
                        duration: 2000,
                        newWindow: true,
                        close: true,
                        gravity: "top",
                        position: "right",
                        stopOnFocus: true,
                    }).showToast();
                }
            });
        });
    }else{
        window.location.href = "http://localhost:8080/"; // Replace with your desired destination URL
    }
})