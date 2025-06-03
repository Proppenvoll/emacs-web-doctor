<script lang="ts">
import '../app.css';
import EmacsDoctorSvg from '$lib/emacs-doctor-svg.svelte';
import type { FormEventHandler, KeyboardEventHandler } from 'svelte/elements';
import { PUBLIC_API_BASE_URL } from '$env/static/public';
import Button from '$lib/button.svelte';

let history: string[] = $state([]);

const response = $derived(
  history.length
    ? history[history.length - 1]
    : 'I am the psychotherapist. Please, describe your problems.',
);

let sliceOfHistoryReversed = $derived(history.slice(0, -1).reverse());

const maxLength = 500;
const answerFormKey = 'answer';
const separator = '\n\n';

let answer = $state('');
const wordCount = $derived(answer.length);

const sendDoctorRequest: FormEventHandler<HTMLFormElement> = async (event) => {
  event.preventDefault();

  const parsedAnswer = answer.replace(/\n$/g, '');

  const answerWithHistory = history.length
    ? history.join(separator) + separator + parsedAnswer
    : parsedAnswer;

  const urlEncodedData = new URLSearchParams({
    [answerFormKey]: answerWithHistory,
  });

  const response = await fetch(`${PUBLIC_API_BASE_URL}/api/doctor`, {
    method: 'post',
    body: urlEncodedData,
  });

  const res = await response.text();

  history = res.split('\n\n').map((element) => element.replaceAll('\n', ' '));
  answer = '';
};

const submitIfNewlines: KeyboardEventHandler<HTMLTextAreaElement> = ({
  key,
  currentTarget,
}) => {
  if (key === 'Enter' && currentTarget.value.at(-1) === '\n') {
    currentTarget.form?.requestSubmit();
  }
};
</script>

<svelte:head>
  <title>Emacs Web Doctor</title>
</svelte:head>

<main
  class="mx-auto flex h-full w-full max-w-(--breakpoint-sm) flex-1 flex-col justify-between p-2"
>
  <div class="*:mx-auto">
    <EmacsDoctorSvg class="max-w-64" />

    <div
      class="relative mt-4 max-w-prose rounded-sm border border-black p-4 before:absolute before:top-0 before:left-1/2 before:-mt-2 before:block before:-translate-x-1/2 before:border-8 before:border-t-0 before:border-transparent before:border-b-black"
    >
      {response}
    </div>

    {#if sliceOfHistoryReversed.length}
      <div class="mt-3 max-w-prose text-neutral-700">
        <div
          class="h-52 space-y-2 rounded-sm mask-[linear-gradient(to_bottom,white_20%,transparent)] text-sm text-neutral-700"
        >
          {#each sliceOfHistoryReversed as historyEntry, index (index)}
            {#if index % 2 === 0}
              <p
                class={historyEntry
                  ? 'max-w-prose'
                  : 'font-mono tracking-wide italic'}
              >
                {historyEntry || 'silence'}
              </p>
            {:else}
              <p class="ml-auto max-w-prose text-right">{historyEntry}</p>
            {/if}
          {/each}
        </div>
      </div>
    {/if}
  </div>

  <form method="POST" onsubmit={sendDoctorRequest}>
    <div class="mx-auto max-w-prose">
      <label class="flex flex-col">
        <span class="flex justify-between text-xs text-neutral-400">
          <span>Your answer</span>
          <span>{wordCount}/{maxLength}</span>
        </span>

        <textarea
          bind:value={answer}
          name={answerFormKey}
          maxlength={maxLength}
          class="mt-1 h-24 resize-none rounded-t border border-neutral-300 bg-neutral-100 p-1 placeholder:text-sm focus:border-purple-800 focus:ring-1 focus:ring-3 focus:ring-purple-800 focus:outline-hidden focus:ring-inset"
          onkeydown={submitIfNewlines}
          placeholder="Pressing two times Enter in a row will send the answer."
        ></textarea>
      </label>

      <Button class="w-full">Submit</Button>
    </div>
  </form>

  <dialog
    open={true}
    class="fixed top-0 m-0 h-dvh w-dvw items-center justify-center bg-black/50 p-2 backdrop-blur-sm open:flex"
  >
    <form
      role="presentation"
      class="w-96 space-y-8 rounded-sm bg-white p-2"
      method="dialog"
    >
      <p>
        Hi there! I am your personal Emacs Therapist. You’ll usually find me in
        my favorite environment: Emacs. I truly enjoy it here, and I’m sure you
        will too. So, go ahead — talk to me, and let's work together to help you
        feel better.
      </p>

      <Button>Lets get started</Button>
    </form>
  </dialog>
</main>
