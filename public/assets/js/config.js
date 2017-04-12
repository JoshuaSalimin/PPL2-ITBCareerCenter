requirejs.config({
    baseUrl: '/public/assets/js',
    paths: {
        jquery: 'jquery.min',
        dropotron: 'jquery.dropotron.min',
        bootstrap: 'bootstrap.min',
        pagination: 'jquery.simplePagination',
        skel: 'skel.min',
        util: 'util',
        fileinput: 'fileinput.min',
        main: 'main'
    },
    shim: {
        'dropotron': ['jquery'],
        'pagination': ['jquery'],
        'util': ['jquery'],
        'fileinput': ['jquery', 'bootstrap'],
        'main': ['jquery', 'skel', 'util']
    }
});