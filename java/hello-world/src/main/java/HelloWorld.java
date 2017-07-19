public class HelloWorld {
    public static String hello(String name) {
        if ("".equals(name) || name == null) {
            name = "World";
        };
        String output = "Hello, " + name + "!";
        return output;
    }
}
