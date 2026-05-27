import router from "kitwork/router";

export const getHello = () => {
    return "Hello from modular helper!";
};

router.get("/modular-route").handle((response) => {
    return response.text("Hello from a route defined in helper.kitwork.js!");
});
