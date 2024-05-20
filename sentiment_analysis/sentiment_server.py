from flask import Flask, request, jsonify
from transformers import pipeline
import logging

app = Flask(__name__)
sentiment_pipeline = pipeline("sentiment-analysis")

logging.basicConfig(level=logging.INFO)

@app.route('/sentiment', methods=['POST'])
def sentiment():
    text = request.json.get('text')
    result = sentiment_pipeline(text)
    return jsonify(result)

if __name__ == '__main__':
    logging.info("Starting sentiment analysis service...")
    app.run(host='0.0.0.0', port=5000)
