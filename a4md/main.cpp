#include <string>
#include <iostream>
#include <fstream>

using namespace::std;

#define OUTNAME "out.html"
#define FILENAME "tmp.html"

int main() {
  ifstream infile;
  infile.exceptions(std::ifstream::badbit);
  try {
    infile.open(FILENAME);
  } catch (const ifstream::failure& e) {
    cerr << "Exception opening/reading file" << endl;
  }

  bool body = false; std::string line;

  std::ofstream ofile;
  ofile.open(OUTNAME); 

  while (getline(infile, line)) {

    if (line == "</body>") {
      body = false;
    }

    if (body) {
      ofile << line << '\n';
    } else if (line == "<body>") {
      body = true;
    }
  }

  ofile.close();
  infile.close();
}
