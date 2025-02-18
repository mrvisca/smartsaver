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
})