// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DO NOT EDIT. GENERATED BY 'gomobile help documentation'.

/*
Gomobile is a tool for building and running mobile apps written in Go.

Installation:

	$ go get golang.org/x/mobile/cmd/gomobile
	$ gomobile init

	Note that until Go 1.5 is released, you must compile Go from tip.

	Clone the source from the tip under $HOME/go directory. On Windows,
	you may like to clone the repo to your user folder, %USERPROFILE%\go.

	  $ git clone https://go.googlesource.com/go $HOME/go

	Go 1.5 requires Go 1.4. Read more about this requirement at
	http://golang.org/s/go15bootstrap.
	Set GOROOT_BOOTSTRAP to the GOROOT of your existing 1.4 installation or
	follow the steps below to checkout go1.4 from the source and build.

	  $ git clone https://go.googlesource.com/go $HOME/go1.4
	  $ cd $HOME/go1.4
	  $ git checkout go1.4.1
	  $ cd src && ./make.bash

	If you clone Go 1.4 to a different destination, set GOROOT_BOOTSTRAP
	environmental variable accordingly.

	Build Go 1.5 and add Go 1.5 bin to your path.

	  $ cd $HOME/go/src && ./make.bash
	  $ export PATH=$PATH:$HOME/go/bin

	Set a GOPATH if no GOPATH is set, add $GOPATH/bin to your path.

	  $ export GOPATH=$HOME
	  $ export PATH=$PATH:$GOPATH/bin

	Now you can get the gomobile tool and initialize.

	  $ go get golang.org/x/mobile/cmd/gomobile
	  $ gomobile init

	It may take a while to initialize gomobile, please wait.

Usage:

	gomobile command [arguments]

Commands:

	bind        build a shared library for android APK and iOS app
	build       compile android APK and iOS app
	init        install android compiler toolchain
	install     compile android APK and iOS app and install on device

Use 'gomobile help [command]' for more information about that command.

NOTE: iOS support is not ready yet.


Build a shared library for android APK and iOS app

Usage:

	gomobile bind [package]

Bind generates language bindings like gobind (golang.org/x/mobile/cmd/gobind)
for a package and builds a shared library for each platform from the go binding
code.

For Android, the bind command produces an AAR (Android ARchive) file that
archives the precompiled Java API stub classes, the compiled shared libraries,
and all asset files in the /assets subdirectory under the package directory.
The output AAR file name is '<package_name>.aar'.

The AAR file is commonly used for binary distribution of an Android library
project and most Android IDEs support AAR import. For example, in Android
Studio (1.2+), an AAR file can be imported using the module import wizard
(File > New > New Module > Import .JAR or .AAR package), and setting it as
a new dependency (File > Project Structure > Dependencies).

This command requires 'javac' (version 1.7+) and Android SDK (API level 9
or newer) to build the library for Android. The environment variable
ANDROID_HOME must be set to the path to Android SDK.

The -v flag provides verbose output, including the list of packages built.

These build flags are shared by the build command.
For documentation, see 'go help build':
	-a
	-i
	-n
	-x
	-tags 'tag list'


Compile android APK and iOS app

Usage:

	gomobile build [-o output] [-i] [build flags] [package]

Build compiles and encodes the app named by the import path.

The named package must define a main function.

If an AndroidManifest.xml is defined in the package directory, it is
added to the APK file. Otherwise, a default manifest is generated.

If the package directory contains an assets subdirectory, its contents
are copied into the APK file.

The -o flag specifies the output file name. If not specified, the
output file name depends on the package built. The output file must end
in '.apk'.

The -v flag provides verbose output, including the list of packages built.

These build flags are shared by the build, install, and test commands.
For documentation, see 'go help build':
	-a
	-i
	-n
	-x
	-tags 'tag list'


Install android compiler toolchain

Usage:

	gomobile init [-u]

Init downloads and installs the Android C++ compiler toolchain.

The toolchain is installed in $GOPATH/pkg/gomobile.
If the Android C++ compiler toolchain already exists in the path,
it skips download and uses the existing toolchain.

The -u option forces download and installation of the new toolchain
even when the toolchain exists.


Compile android APK and iOS app and install on device

Usage:

	gomobile install [package]

Install compiles and installs the app named by the import path on the
attached mobile device.

This command requires the 'adb' tool on the PATH.

See the build command help for common flags and common behavior.
*/
package main