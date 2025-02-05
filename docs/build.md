## Building Bisq

1. **Install Git LFS**

   Bisq uses Git LFS (Large File Storage) to track certain large binary files. Follow the instructions at https://git-lfs.github.com to install it, then run the following to command to verify the installation:

   ```sh
   git lfs version
   ```

   You should see the version of Git LFS you installed, for example:

   ```sh
   git-lfs/2.10.0 (GitHub; darwin amd64; go 1.13.6)
   ```

2. **Clone Bisq**

   ```sh
   git clone https://github.com/bisq-network/bisq
   cd bisq
   ```

3. **Pull LFS data**

   ```sh
   git lfs pull
   ```

4. **Build Bisq**

   On macOS and Linux, execute:
   ```sh
   ./gradlew build
   ```

   On Windows:
   ```cmd
   gradlew.bat build
   ```

### Important notes

1. You do _not_ need to install Gradle to build Bisq. The `gradlew` shell script will install it for you, if necessary.

2. Bisq currently works with JDK 11 only. JDK 12 and above are not supported. You can find out which
   version you have with:

   ```sh
   javac -version
   ```

If you do not have JDK 11 installed, check out scripts in the [scripts](../scripts) directory or download it manually from https://jdk.java.net/archive/.

## Running Bisq

Once Bisq is installed, its executables will be available in the root project directory. Run **Bisq Desktop** as follows:

On macOS and Linux:
```sh
./bisq-desktop
```

On Windows:
```cmd
bisq-desktop.bat
```

## See also

 - [Importing Bisq into IntelliJ IDEA](./idea-import.md)
 - [Bisq development environment setup guide](./dev-setup.md)
