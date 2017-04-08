requirejs.config({
    baseUrl: '/public/assets/js',
    paths: {
        jquery: 'jquery.min',
        dropotron: 'jquery.dropotron.min',
        pagination: 'jquery.simplePagination',
        skel: 'skel.min',
        util: 'util',
        main: 'main'
    },
    shim: {
        'dropotron': ['jquery'],
        'pagination': ['jquery'],
        'util': ['jquery'],
        'main': ['jquery', 'skel', 'util']
    }
});