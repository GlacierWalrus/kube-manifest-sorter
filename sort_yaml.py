import yaml
from ruamel.yaml import YAML
from operator import itemgetter
import sys

def load_documents(filename):
    with open(filename, 'r') as stream:
        return list(yaml.load_all(stream, Loader=yaml.FullLoader))

def sort_documents(documents):
    return sorted(documents, key=lambda doc: (doc['apiVersion'], doc['kind'], doc['metadata']['name']))

def write_documents(filename, sorted_documents):
    ryaml = YAML()
    ryaml.explicit_start = True
    with open(filename, 'w') as yaml_file:
        for doc in sorted_documents:
            ryaml.dump(doc, yaml_file)

def main():
    if len(sys.argv) != 3:
        print("Usage: python sort_yaml.py <input_file> <output_file>")
        return

    input_filename = sys.argv[1]
    output_filename = sys.argv[2]
    documents = load_documents(input_filename)
    sorted_documents = sort_documents(documents)
    write_documents(output_filename, sorted_documents)

main()
