#!/usr/bin/env python3
from ruamel.yaml import YAML
from operator import itemgetter
import sys

def load_documents(filename):
    ryaml = YAML()
    with open(filename, 'r') as stream:
        return list(ryaml.load_all(stream))

def sort_documents(documents):
    return sorted(documents, key=lambda doc: (doc['apiVersion'], doc['kind'], doc['metadata']['name']))

def write_documents(sorted_documents):
    ryaml = YAML()
    ryaml.explicit_start = True
    for doc in sorted_documents:
        ryaml.dump(doc, sys.stdout)

def main():
    if len(sys.argv) != 2:
        print("Usage: python sort_yaml.py <input_file>")
        return

    input_filename = sys.argv[1]
    documents = load_documents(input_filename)
    sorted_documents = sort_documents(documents)
    write_documents(sorted_documents)

main()
