const { router, log, render, http, database, go } = kitwork();


router.get("/hello").handle((response) => {
    return response.text("hello world");
});

router.get("/").handle((response) => {
    return response.text("Welcome to Kitwork");
});

router.get("/teststatic").static("1m").handle((response) => {
    return response.text("this is static cached text at " + new Date().toISOString());
});