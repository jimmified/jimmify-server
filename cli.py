import requests
import json

class API():

	def __init__(self, remote):
		if remote:
			self.url = "http://shibboleth.student.rit.edu/"
		else:
			self.url = "http://localhost:3000/api/"

	def query(self,text,typ):
		r = {}
		r['text'] = text
		r['type'] = typ
		r = requests.post(self.url + "query", data=json.dumps(r), verify=False)
		data = r.json()
		print(data)

	def charge(self, charge, query):
		r = {}
		r['charge'] = charge
		r['query'] = query
		r = requests.post(self.url + "charge", data=json.dumps(r),
		verify=False)
		data = r.json()
		print(data)

	def queue(self):
		r = requests.get(self.url + "queue", verify=False)
		data = r.json()
		print(data)

	def answer(self, key, answer):
		r = {}
		r['key'] = key
		r['answer'] = answer
		r = requests.post(self.url + "answer", data=json.dumps(r), verify=False)
		data = r.json()
		print(data)

	def check(self, key):
		r = {}
		r['key'] = key
		r = requests.post(self.url + "check", data=json.dumps(r), verify=False)
		data = r.json()
		print(data)

	def recent(self):
		r = requests.get(self.url + "recent", verify=False)
		data = r.json()
		print(data)

if __name__ == '__main__':
	print("Jimmy CLI Starting.. ")
	i = input("local or remote? ")
	if i == "remote":
		remote = True
	else:
		remote = False
	jimmy = API(remote)
	#Start REPL
	print("Use command 'help' for more options.")
	while True:
		i = input("Enter command> ")
		if i == "query":
			text = input("query> ")
			typ = "search"
			jimmy.query(text, typ)
		elif i == "queue":
			jimmy.queue()
		elif i == "answer":
			key = int(input("key> "))
			answer = input("answer> ")
			jimmy.answer(key, answer)
		elif i == "check":
			key = int(input("key> "))
			jimmy.check(key)
		elif i == "recent":
			jimmy.recent()
		elif i == "help":
			print("Avaliable commands: help, query, queue, answer, check, recent")
		elif i == "quit":
			break
