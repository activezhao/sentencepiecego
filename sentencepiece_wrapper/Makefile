CXX=g++
INCLUDES=-I. -I../src/ -I../
CXXFLAGS=-fPIC -pthread -Wall -std=c++17 -O2 -march=x86-64 $(INCLUDES)
LDFLAGS=-shared
OBJS=sentencepiece_wrapper.o
TARGET=libsentencepiecego.so
SRC=src/CMakeFiles/sentencepiece.dir

all: $(TARGET)

$(OBJS): sentencepiece_wrapper.h sentencepiece_wrapper.cc
	$(CXX) $(CXXFLAGS) -c sentencepiece_wrapper.cc

$(TARGET): $(OBJS)
	$(CXX) $(LDFLAGS) -o $(TARGET) $(OBJS) src/libsentencepiece.a
