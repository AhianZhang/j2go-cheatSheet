package syntax;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

class IoDemo {
    public static void main(String[] args) throws IOException {
        String filePath = "path/j2go-cheatSheet/java/syntax/Main.java";
        byte[] bytes = Files.readAllBytes(Paths.get(filePath));
        System.out.println(new String(bytes));
    }

}