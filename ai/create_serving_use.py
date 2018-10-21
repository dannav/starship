# create a model to serve as API with tensorflow/serving from the universal-sentence-encoder
# this api will be used to generate word embeddings

import tensorflow as tf
import tensorflow_hub as hub
from tensorflow.saved_model import simple_save

export_dir = "./tfserving/use/00000001"
with tf.Session(graph=tf.Graph()) as sess:
    module = hub.Module("https://tfhub.dev/google/universal-sentence-encoder/2")
    input_params = module.get_input_info_dict()
    # take a look at what tensor does the model accepts - 'text' is input tensor name

    text_input = tf.placeholder(name='text', dtype=input_params['text'].dtype,
        shape=input_params['text'].get_shape())
    sess.run([tf.global_variables_initializer(), tf.tables_initializer()])

    embeddings = module(text_input)

    simple_save(sess,
        export_dir,
        inputs={'text': text_input},
        outputs={'embeddings': embeddings},
        legacy_init_op=tf.tables_initializer())